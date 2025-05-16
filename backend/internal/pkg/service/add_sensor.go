package service

import (
	"fmt"
	"sensor-simulator/internal/pkg/dto"
)

func (s *SimulatorService) AddSensor(dto dto.Simulator) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	simulator, newDependencies, err := s.simulatorFromDTO(dto)
	if err != nil {
		return fmt.Errorf("unable to create simulator. Err: %w", err)
	}

	if _, ok := s.simulators[SimulatorName(simulator.GetName())]; ok {
		return fmt.Errorf("simulator with that name already exists.")
	}

	s.simulators[SimulatorName(simulator.GetName())] = simulator
	for _, dependency := range newDependencies {
		simulatorDependencies, ok := s.dependencies[SimulatorName(dependency)]
		if !ok {
			simulatorDependencies = map[SimulatorName]struct{}{}
		}

		simulatorDependencies[SimulatorName(dependency)] = struct{}{}
	}

	return nil
}
