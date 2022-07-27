package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"pet_care.baroni.tech/src/cache"
	"pet_care.baroni.tech/src/pet_status"
)

func handleStart(w http.ResponseWriter, req *http.Request) {
	petNameBuf, err := ioutil.ReadAll(req.Body)
	if err != nil || len(petNameBuf) < 2 {
		fmt.Fprintf(w, "Missing pet name (POST body)")
		return
	}
	petName := string(petNameBuf[:])
	serviceId := uuid.Must(uuid.NewRandom()).String()

	cache.SetString(serviceId+"_Name", string(petName))
	cache.SetString(serviceId+"_Status", pet_status.Received.String())

	fmt.Fprintf(w, "Service created %s for pet %s", serviceId, petName)
}

func handleStatus(w http.ResponseWriter, req *http.Request) {
	serviceIdBuf, err := ioutil.ReadAll(req.Body)
	if err != nil || len(serviceIdBuf) < 2 {
		fmt.Fprintf(w, "Missing serviceId (POST body)")
		return
	}
	serviceId := string(serviceIdBuf[:])
	petName := cache.GetString(serviceId + "_Name")
	status := cache.GetString(serviceId + "_Status")
	if len(petName) == 0 || len(status) == 0 {
		fmt.Fprintf(w, "serviceId not found")
		return
	}

	fmt.Fprintf(w, "Service for %s for pet %s is at %s", serviceId, petName, status)
}

func main() {
	http.HandleFunc("/api/pet/start", handleStart)
	http.HandleFunc("/api/pet/status", handleStatus)

	http.ListenAndServe(":8080", nil)
}
