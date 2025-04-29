package service

import (
	"context"
	"fmt"
	pb "sensor-simulator/gen/sensor_simulator/proto/simulator"
)

func (s *SimulatorService) DeleteSensor(ctx context.Context, req *pb.DeleteSensorRequest) (*pb.DeleteSensorResponse, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.simulators[SimulatorName(req.GetName())]; !ok {
		return nil, fmt.Errorf("simulator with that name not exist.")
	}

	simulatorDependencies, ok := s.dependencies[SimulatorName(req.GetName())]
	if ok {
		if len(simulatorDependencies) != 0 {
			return nil, fmt.Errorf("simulators %v dependant from this simulator.", simulatorDependencies)
		}
	}

	delete(s.dependencies, SimulatorName(req.GetName()))

	for _, v := range s.dependencies {
		if _, ok := v[SimulatorName(req.GetName())]; ok {
			delete(v, SimulatorName(req.GetName()))
		}
	}

	return &pb.DeleteSensorResponse{}, nil
}
