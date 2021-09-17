package factory

import (
	".main.go/vehicle"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {
	s.adapter = New()
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSamble() {
	cars := 8
	wg := sync.WaitGroup{}
	wg.Add(cars)
	results := make(chan *vehicle.Car, cars)
	s.adapter.FinishedCars = results
	go func() {
		for car := range results {
			s.Assert().NotNil(car)
			if car != nil {
				s.Assert().Equal("Assembled", car.Engine)
				s.Assert().Equal("Assembled", car.Chassis)
				s.Assert().Equal("Assembled", car.Dash)
				s.Assert().Equal("Assembled", car.Electronics)
				s.Assert().Equal("Assembled", car.Sits)
				s.Assert().Equal("Assembled", car.Tires)
				s.Assert().Equal("Assembled", car.Windows)
			}
			wg.Done()
		}
	}()
	s.adapter.StartAssemblingProcess(cars)

	wg.Wait()
}
