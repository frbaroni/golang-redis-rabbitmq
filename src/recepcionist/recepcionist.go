package recepcionist

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"petcare.baroni.tech/src/animal_service"
)

func handleStart(w http.ResponseWriter, req *http.Request) {
	petNameBuf, err := ioutil.ReadAll(req.Body)
	if err != nil || len(petNameBuf) < 2 {
		fmt.Fprintf(w, "Missing pet name (POST body)")
		return
	}
	service := animal_service.CreateService(string(petNameBuf[:]))
	res := fmt.Sprintf("Service created %s for pet %s", service.ServiceId, service.Name)
	fmt.Println(res)
	fmt.Fprint(w, res)
}

func handleStatus(w http.ResponseWriter, req *http.Request) {
	serviceIdBuf, err := ioutil.ReadAll(req.Body)
	if err != nil || len(serviceIdBuf) < 2 {
		fmt.Fprintf(w, "Missing serviceId (POST body)")
		return
	}
	serviceId := string(serviceIdBuf[:])
	service, err := animal_service.Load(serviceId)
	if err != nil {
		fmt.Fprintf(w, "Service not found")
		return
	}

	res := fmt.Sprintf("Service status for %s for pet %s is at %s", serviceId, service.Name, service.Status)
	fmt.Println(w, res)
	fmt.Fprint(w, res)
}

func RunRecepcionist() {
	fmt.Println("Looking for Pets to take care of on port :8080")

	http.HandleFunc("/api/pet/start", handleStart)
	http.HandleFunc("/api/pet/status", handleStatus)

	http.ListenAndServe(":8080", nil)
}
