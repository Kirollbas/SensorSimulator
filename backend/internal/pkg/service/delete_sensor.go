package service

import (
	"fmt"
)

func (s *SimulatorService) DeleteSensor(name string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.simulators[SimulatorName(name)]; !ok {
		return fmt.Errorf("simulator with that name not exist.")
	}

	simulatorDependencies, ok := s.dependencies[SimulatorName(name)]
	if ok {
		if len(simulatorDependencies) != 0 {
			return fmt.Errorf("simulators %v dependant from this simulator.", simulatorDependencies)
		}
	}

	delete(s.dependencies, SimulatorName(name))

	for _, v := range s.dependencies {
		if _, ok := v[SimulatorName(name)]; ok {
			delete(v, SimulatorName(name))
		}
	}

	delete(s.simulators, SimulatorName(name))

	return nil
}
