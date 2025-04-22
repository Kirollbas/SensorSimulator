package modifier

import (
	"sensor-simulator/internal/pkg/domain/simulator"
)

type ProgressingOffset struct {
	offsetChange float64
	offset       float64

	currentTick   uint64
	ticksToChange uint64
}

func NewProgressingOffsetModifier(
	offsetChange float64,
	ticksToChange uint64,

) (*ProgressingOffset, error) {
	return &ProgressingOffset{
		offsetChange:  offsetChange,
		offset:        0,
		currentTick:   0,
		ticksToChange: ticksToChange,
	}, nil
}

func (po *ProgressingOffset) ApplyModifier(point simulator.PointState) simulator.PointState {
	if po.currentTick >= po.ticksToChange {
		po.currentTick = 0
		po.offset += po.offsetChange
	}

	point.Value += po.offset
	po.currentTick++

	return point
}
