package service

import "sensor-simulator/internal/pkg/measurement/simulator"

type SimulatorName string

type SensorService struct {
	simulators   map[SimulatorName]simulator.Simulator
	dependencies map[SimulatorName][]SimulatorName
}
