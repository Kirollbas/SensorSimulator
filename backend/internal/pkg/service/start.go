package service

import (
	"context"
	"fmt"
	pb "sensor-simulator/gen/sensor_simulator/proto/simulator"
	"sensor-simulator/internal/pkg/domain/simulator"
)

func (s *SimulatorService) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	simulators := []*simulator.Simulator{}
	for _, simulator := range s.simulators {
		simulators = append(simulators, simulator)
	}

	err := s.opcuaServer.Start(ctx, simulators)
	if err != nil {
		return nil, fmt.Errorf("unable to launch OPC UA server. Err: %w", err)
	}

	err = s.modbusServer.Start(ctx, simulators)
	if err != nil {
		return nil, fmt.Errorf("unable to launch modbus server. Err: %w", err)
	}

	for _, simulator := range simulators {
		simulator.Start()
	}

	return &pb.StartResponse{}, nil
}
