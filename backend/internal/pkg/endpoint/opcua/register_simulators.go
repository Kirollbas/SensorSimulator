package opcua

import (
	"sensor-simulator/internal/pkg/domain/simulator"
	"sensor-simulator/internal/pkg/endpoint/opcua/node"

	"github.com/gopcua/opcua/id"
)

func (s *Server) registerSimulators(simulators []*simulator.Simulator) {
	for _, simulator := range simulators {
		name := simulator.GetName()
		nns_obj := s.commonNamespace.Objects()

		base := s.commonNamespace.AddNewVariableNode(name+"_base", float64(0))
		nns_obj.AddRef(base, id.HasComponent, true)

		modified := s.commonNamespace.AddNewVariableNode(name, float64(0))
		nns_obj.AddRef(modified, id.HasComponent, true)

		simulator.AddEndpointUpdater(node.NewNode(
			s.commonNamespace,
			base,
			modified,
		))
	}
}
