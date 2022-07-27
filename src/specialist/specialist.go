package specialist

import (
	"fmt"
	"math/rand"
	"time"

	"petcare.baroni.tech/src/animal_service"
	"petcare.baroni.tech/src/messaging"
	"petcare.baroni.tech/src/pet_job"
)

func RunSpecialist() {
	fmt.Println("Looking for Pets to service on RabbitMQ")
	err := messaging.Consume(func(serviceId string) {
		service, err := animal_service.Load(serviceId)
		if err != nil {
			fmt.Println("Service not found " + serviceId)
			return
		}

		var next pet_job.PetJob
		current := pet_job.Parse(service.Status)
		switch current {
		case pet_job.Received:
			next = pet_job.Brushed
		case pet_job.Brushed:
			next = pet_job.Washed
		case pet_job.Washed:
			next = pet_job.Dried
		case pet_job.Dried:
			next = pet_job.Complete
		default:
			next = pet_job.Complete
		}

		fmt.Println("Caring for service " + service.String())
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		if next != current {
			animal_service.UpdateStatus(&service, next)
		}
	})
	if err != nil {
		fmt.Println(err)
	}
}
