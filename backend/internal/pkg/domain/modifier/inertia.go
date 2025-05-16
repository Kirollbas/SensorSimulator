package modifier

import (
	"math"
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
	"time"
)

type Inertia struct {
	value      float64
	period     time.Duration
	startValue float64

	maxChangePerTick float64
	currentValue     float64
}

func NewInertiaModifier(
	value float64,
	tickPeriod time.Duration,
	period time.Duration,
	startValue float64,
) (*Inertia, error) {
	speed := value * float64(tickPeriod) / float64(period)

	return &Inertia{
		value:            value,
		period:           period,
		startValue:       startValue,
		maxChangePerTick: speed,
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

func (i *Inertia) ToDTO() dto.Modifier {
	return dto.Modifier{
		Type: dto.ModifierTypeInertia,
		Data: dto.InertitaModifier{
			Value: i.value,
			Period: dto.Duration{
				Duration: i.period,
			},
		},
	}
}
