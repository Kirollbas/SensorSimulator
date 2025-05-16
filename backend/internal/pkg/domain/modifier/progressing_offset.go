package modifier

import (
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
	"time"
)

type ProgressingOffset struct {
	interval     time.Duration
	offsetChange float64

	offset float64

	currentTick   uint64
	ticksToChange uint64
}

func NewProgressingOffsetModifier(
	offsetChange float64,
	interval time.Duration,
	tickPeriod time.Duration,
) (*ProgressingOffset, error) {
	ticksToChange := uint64(interval / tickPeriod)

	return &ProgressingOffset{
		interval:      interval,
		offsetChange:  offsetChange,
		offset:        0,
		currentTick:   0,
		ticksToChange: ticksToChange,
	}, nil
}

func (po *ProgressingOffset) Restart() {
	po.offset = 0
	po.currentTick = 0
}

func (po *ProgressingOffset) ApplyModifier(point state.PointState) state.PointState {
	if po.currentTick >= po.ticksToChange {
		po.currentTick = 0
		po.offset += po.offsetChange
	}

	point.Value += po.offset
	po.currentTick++

	return point
}

func (po *ProgressingOffset) ToDTO() dto.Modifier {
	return dto.Modifier{
		Type: dto.ModifierTypeProgressingOffset,
		Data: dto.ProgressingOffsetModifier{
			Value: po.offsetChange,
			Interval: dto.Duration{
				Duration: po.interval,
			},
		},
	}
}
