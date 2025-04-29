package modbus

import (
	"context"
	"fmt"
)

func (s *Server) Start(ctx context.Context, simulators []Simulator) error {
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
