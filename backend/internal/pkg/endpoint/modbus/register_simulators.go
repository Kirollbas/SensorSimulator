package modbus

import (
	"fmt"
	"math"
	"sensor-simulator/internal/pkg/domain/simulator"
	"sensor-simulator/internal/pkg/endpoint/modbus/node"
)

func (s *Server) registerSimulators(simulators []*simulator.Simulator) error {
	checkMap := make(map[int]struct{})

	for _, simulator := range simulators {
		address := simulator.GetAddress()

		if int(address)+8 > math.MaxInt16 {
			return fmt.Errorf("unable to add simulator on address %d, not enough registers", address)
		}

		for i := 0; i < 8; i++ {
			if _, ok := checkMap[int(address)+i]; ok {
				return fmt.Errorf("unable to add simulator on address %d, it crosses with other simulator", address)
			}

			checkMap[int(address)+i] = struct{}{}
		}
	}

	for _, simulator := range simulators {
		node := node.NewSimulatorNode()

		simulator.AddEndpointUpdater(node)
		err := s.handler.AddNode(simulator.GetAddress(), node)
		if err != nil {
			return fmt.Errorf("unable to add node. Err: %w", err)
		}
	}

	return nil
}
