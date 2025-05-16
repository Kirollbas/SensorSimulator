package service

import "sensor-simulator/internal/pkg/dto"

func (s *SimulatorService) GetSimulatorDtos() []dto.SimulatorWithStatus {
	s.mx.Lock()
	defer s.mx.Unlock()

	dtos := []dto.SimulatorWithStatus{}
	for _, simulator := range s.simulators {
		dtos = append(dtos, dto.SimulatorWithStatus{
			Simulator: simulator.ToDTO(),
			IsActive:  simulator.IsActive(),
		})
	}

	return dtos
}
