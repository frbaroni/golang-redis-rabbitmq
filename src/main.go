package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"petcare.baroni.tech/src/recepcionist"
	"petcare.baroni.tech/src/specialist"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	runRecepcionist := flag.Bool("recepcionist", false, "Start as a server that handles HTTP /api/pet requests")
	runSpecialist := flag.Bool("specialist", false, "Start as a RabbitMQ handler to take care of the pets")
	flag.Parse()

	if *runRecepcionist && !*runSpecialist {
		recepcionist.RunRecepcionist()
	} else if !*runRecepcionist && *runSpecialist {
		specialist.RunSpecialist()
	} else {
		log.Panicln("Please specify --recepcionist XOR --specialist")
	}
}
