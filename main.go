package main

import (
	".main.go/factory"
	".main.go/vehicle"
	"log"
	"sync"
)

const carsAmount = 100

func main() {
	output := make(chan *vehicle.Car, carsAmount)

	f := factory.New()
	f.FinishedCars = output

	wg := sync.WaitGroup{}
	wg.Add(carsAmount)
	go func() {
		for car := range output {
			log.Println(car.AssembleLog)
			log.Println(car.TestingLog)
			wg.Done()
		}
	}()

	f.StartAssemblingProcess(carsAmount)
	wg.Wait()
}
