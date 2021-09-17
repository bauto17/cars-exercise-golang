package factory

import (
	".main.go/assemblyspot"
	".main.go/vehicle"
	"log"
)

const assemblySpots int = 5

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
	FinishedCars    chan *vehicle.Car
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
		FinishedCars:    nil,
	}

	totalAssemblySpots := 0

	for {
		factory.AssemblingSpots <- assemblyspot.NewAssemblySpot(totalAssemblySpots)

		totalAssemblySpots++

		if totalAssemblySpots >= assemblySpots {
			break
		}
	}

	return factory
}

func (f *Factory) StartAssemblingProcess(amountOfVehicles int) {
	vehicleList := f.generateVehicleLots(amountOfVehicles)

	for _, car := range vehicleList {

		idleSpot := <-f.AssemblingSpots
		idleSpot.SetVehicle(car)

		go func() {
			log.Println("Assembling vehicle...")
			result, err := idleSpot.AssembleVehicle()
			if err == nil {
				result.TestingLog = f.testCar(result)
				result.AssembleLog = idleSpot.GetAssembledLogs()
			}

			idleSpot.CleanSpot()
			f.FinishedCars <- result
			f.AssemblingSpots <- idleSpot
		}()
	}
}

func (Factory) generateVehicleLots(amountOfVehicles int) []vehicle.Car {
	var vehicles = []vehicle.Car{}
	var index = 0

	for {
		vehicles = append(vehicles, vehicle.Car{
			Id:            index,
			Chassis:       "NotSet",
			Tires:         "NotSet",
			Engine:        "NotSet",
			Electronics:   "NotSet",
			Dash:          "NotSet",
			Sits:          "NotSet",
			Windows:       "NotSet",
			EngineStarted: false,
		})

		index++

		if index >= amountOfVehicles {
			break
		}
	}

	return vehicles
}

func (f *Factory) testCar(car *vehicle.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
