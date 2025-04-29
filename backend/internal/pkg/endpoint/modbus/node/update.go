package node

import (
	"math"
	"sensor-simulator/internal/pkg/domain/state"
	"sync/atomic"
)

func (n *SimulatorNode) Update(state state.PointState) {
	atomic.StoreUint64(&n.baseValue, math.Float64bits(state.BaseValue))
	atomic.StoreUint64(&n.modifiedValue, math.Float64bits(state.Value))
}
