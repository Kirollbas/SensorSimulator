package service

import (
	"context"
	"fmt"
	pb "sensor-simulator/gen/sensor_simulator/proto/simulator"
)

func (s *SimulatorService) AddSensor(ctx context.Context, req *pb.AddSensorRequest) (*pb.AddSensorResponse, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	simulator, newDependencies, err := s.simulatorFromPb(req.GetSimulator())
	if err != nil {
		return nil, fmt.Errorf("unable to create simulator. Err: %w", err)
	}

	if _, ok := s.simulators[SimulatorName(simulator.GetName())]; ok {
		return nil, fmt.Errorf("simulator with that name already exists.")
	}

	s.simulators[SimulatorName(simulator.GetName())] = simulator
	for _, dependency := range newDependencies {
		simulatorDependencies, ok := s.dependencies[SimulatorName(dependency)]
		if !ok {
			simulatorDependencies = map[SimulatorName]struct{}{}
		}

		simulatorDependencies[SimulatorName(dependency)] = struct{}{}
	}

	return &pb.AddSensorResponse{}, nil
}
