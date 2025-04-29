package modifier

import "sensor-simulator/internal/pkg/domain/simulator"

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

func (o *ConstantOffset) ApplyModifier(point simulator.PointState) simulator.PointState {
	point.Value += o.offset
	return point
}
