package opcua

import (
	"context"
	"fmt"
)

func (s *Server) Start(ctx context.Context, simulators []Simulator) {
	s.registerSimulators(simulators)

	if err := s.server.Start(ctx); err != nil {
		fmt.Printf("Error starting server, exiting: %s", err)
	}
}
