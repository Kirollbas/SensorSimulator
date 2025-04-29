package node

import (
	"math"
	"sensor-simulator/internal/pkg/domain/simulator"
	"sync/atomic"
)

func (n *SimulatorNode) Update(state simulator.PointState) {
	atomic.StoreUint64(&n.baseValue, math.Float64bits(state.BaseValue))
	atomic.StoreUint64(&n.modifiedValue, math.Float64bits(state.Value))
}
