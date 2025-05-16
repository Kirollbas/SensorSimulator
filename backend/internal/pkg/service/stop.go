package service

import (
	"context"
	"fmt"
)

func (s *SimulatorService) Stop(ctx context.Context) error {
	for _, simulator := range s.simulators {
		simulator.Finish()
		simulator.Restart()
	}

	err := s.opcuaServer.Close()
	if err != nil {
		return fmt.Errorf("unable to stop OPC UA server. Err: %w", err)
	}

	err = s.modbusServer.Close()
	if err != nil {
		return fmt.Errorf("unable to stop modbus server. Err: %w", err)
	}

	return nil
}
