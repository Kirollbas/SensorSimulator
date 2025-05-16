package service

import (
	"context"
	"fmt"
	"sensor-simulator/internal/pkg/domain/simulator"
)

func (s *SimulatorService) Start(ctx context.Context) error {
	simulators := []*simulator.Simulator{}
	for _, simulator := range s.simulators {
		simulators = append(simulators, simulator)
	}

	err := s.opcuaServer.Start(ctx, simulators)
	if err != nil {
		return fmt.Errorf("unable to launch OPC UA server. Err: %w", err)
	}

	err = s.modbusServer.Start(ctx, simulators)
	if err != nil {
		return fmt.Errorf("unable to launch modbus server. Err: %w", err)
	}

	for _, simulator := range simulators {
		simulator.Start()
	}

	return nil
}
