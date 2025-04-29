package modbus

import (
	"context"
	"fmt"
	"sensor-simulator/internal/pkg/domain/simulator"
)

func (s *Server) Start(ctx context.Context, simulators []*simulator.Simulator) error {
	err := s.registerSimulators(simulators)
	if err != nil {
		return fmt.Errorf("unable to register simulators. Err: %w", err)
	}

	err = s.server.Start()
	if err != nil {
		return fmt.Errorf("unable to start modbus server. Err: %w", err)
	}

	return nil
}
