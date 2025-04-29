package modifier

import "sensor-simulator/internal/pkg/domain/state"

type ConstantOffset struct {
	offset float64
}

func NewConstantOffsetModifier(
	offset float64,
) (*ConstantOffset, error) {
	return &ConstantOffset{
		offset: offset,
	}, nil
}

func (o *ConstantOffset) Restart() {}

func (o *ConstantOffset) ApplyModifier(point state.PointState) state.PointState {
	point.Value += o.offset
	return point
}
