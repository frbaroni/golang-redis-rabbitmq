package animal_service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"petcare.baroni.tech/src/cache"
	"petcare.baroni.tech/src/messaging"
	"petcare.baroni.tech/src/pet_job"
)

type JobService struct {
	Job  string
	Date string
}

type AnimalService struct {
	Name      string
	ServiceId string
	Status    string
	History   []JobService
}

func Parse(buf string) (AnimalService, error) {
	var service AnimalService
	err := json.Unmarshal([]byte(buf), &service)
	return service, err
}

func (s AnimalService) String() string {
	buf, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func (s AnimalService) Save() AnimalService {
	cache.SetString("petcare:"+s.ServiceId, s.String())
	messaging.Emit(s.ServiceId)
	return s
}

func Load(serviceId string) (AnimalService, error) {
	return Parse(cache.GetString("petcare:" + serviceId))
}

func UpdateStatus(s *AnimalService, status pet_job.PetJob) AnimalService {
	s.Status = status.String()
	s.History = append(s.History, JobService{
		Job:  status.String(),
		Date: time.Now().String(),
	})
	return s.Save()
}

func CreateService(petName string) AnimalService {
	service := AnimalService{
		Name:      petName,
		ServiceId: uuid.Must(uuid.NewRandom()).String(),
		Status:    pet_job.Received.String(),
		History:   make([]JobService, 0),
	}
	return UpdateStatus(&service, pet_job.Received)
}
