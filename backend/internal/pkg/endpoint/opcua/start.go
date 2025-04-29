package opcua

import (
	"context"
	"fmt"
	"sensor-simulator/internal/pkg/domain/simulator"
)

func (s *Server) Start(ctx context.Context, simulators []*simulator.Simulator) error {
	s.registerSimulators(simulators)

	if err := s.server.Start(ctx); err != nil {
		return fmt.Errorf("unable to start opcua server. Err:%w", err)
	}

	return nil
}
