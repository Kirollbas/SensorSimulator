package modifier

import (
	"math"
	"sensor-simulator/internal/pkg/domain/state"
)

type Inertia struct {
	startValue       float64
	maxChangePerTick float64
	currentValue     float64
}

func NewInertiaModifier(
	maxChangePerTick float64,
	startValue float64,
) (*Inertia, error) {
	return &Inertia{
		startValue:       startValue,
		maxChangePerTick: maxChangePerTick,
		currentValue:     startValue,
	}, nil
}

func (i *Inertia) Restart() {
	i.currentValue = i.startValue
}

func (i *Inertia) ApplyModifier(point state.PointState) state.PointState {
	diff := point.Value - i.currentValue

	if math.Abs(diff) > i.maxChangePerTick {
		diff = math.Copysign(i.maxChangePerTick, diff)
	}

	i.currentValue += diff
	point.Value = i.currentValue

	return point
}
