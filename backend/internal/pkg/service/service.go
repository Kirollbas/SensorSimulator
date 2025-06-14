package service

import (
	"context"
	"sensor-simulator/internal/pkg/domain/simulator"
	"sync"
)

type SimulatorName string

type endpointServer interface {
	Start(ctx context.Context, simulators []*simulator.Simulator) error
	Close() error
}

type SimulatorService struct {
	modbusServer endpointServer
	opcuaServer  endpointServer

	simulators   map[SimulatorName]*simulator.Simulator
	dependencies map[SimulatorName]map[SimulatorName]struct{}

	mx sync.Mutex
}

func NewSimulatorService(
	modbusServer endpointServer,
	opcuaServer endpointServer,
) *SimulatorService {
	return &SimulatorService{
		modbusServer: modbusServer,
		opcuaServer:  opcuaServer,
		simulators:   map[SimulatorName]*simulator.Simulator{},
		dependencies: map[SimulatorName]map[SimulatorName]struct{}{},
	}
}
