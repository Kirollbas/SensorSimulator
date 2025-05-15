package opcua

import (
	"fmt"
)

func (s *Server) Close() error {
	err := s.server.Close()
	if err != nil {
		return fmt.Errorf("unable to stop opcua server. Err:%w", err)
	}

	s.resetServer()

	return nil
}
