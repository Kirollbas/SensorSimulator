package modbus

import "fmt"

func (s *Server) Close() error {
	err := s.server.Stop()
	if err != nil {
		return fmt.Errorf("unable to stop modbus server. Err: %w", err)
	}

	s.handler.ClearNodes()

	return nil
}
