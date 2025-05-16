package modifier

import (
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
)

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

func (o *ConstantOffset) ToDTO() dto.Modifier {
	return dto.Modifier{
		Type: dto.ModifierTypeConstantOffset,
		Data: dto.ConstantOffsetModifier{
			Offset: o.offset,
		},
	}
}
