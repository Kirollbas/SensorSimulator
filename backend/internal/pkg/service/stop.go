package service

import (
	"context"
	"fmt"
	pb "sensor-simulator/gen/sensor_simulator/proto/simulator"
)

func (s *SimulatorService) Stop(ctx context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	for _, simulator := range s.simulators {
		simulator.Finish()
		simulator.Restart()
	}

	err := s.opcuaServer.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to stop OPC UA server. Err: %w", err)
	}

	err = s.modbusServer.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to stop modbus server. Err: %w", err)
	}

	return &pb.StopResponse{}, nil
}
