package modifier

import (
	"math"
	"sensor-simulator/internal/pkg/domain/state"
)

type Hysteresis struct {
	percentage float64

	maxHysteresis float64
	ticksCenter   uint64
}

func NewHysteresisModifier(
	percentage uint64,
) (*Hysteresis, error) {
	return &Hysteresis{
		percentage: float64(percentage) / 100,
	}, nil
}

func (h *Hysteresis) Restart() {
	h.maxHysteresis = 0
	h.ticksCenter = 0
}

func (h *Hysteresis) UpdateState(state state.SimulatorBaseState) {
	h.maxHysteresis = h.percentage * (state.NextPoint - state.PreviousPoint)
	h.ticksCenter = state.TicksDistance / 2
}

func (h *Hysteresis) ApplyModifier(point state.PointState) state.PointState {
	point.Value += h.maxHysteresis *
		(1 - math.Pow(math.Abs(float64(h.ticksCenter)-float64(point.Tick))/float64(h.ticksCenter), 2))

	return point
}
