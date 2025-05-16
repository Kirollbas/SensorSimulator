package base

import (
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
)

type ConstantSimulator struct {
	value float64
}

func NewConstantSimulator(
	value float64,
) *ConstantSimulator {
	return &ConstantSimulator{
		value: value,
	}
}

func (c *ConstantSimulator) Iterate() state.PointState {
	return state.PointState{
		BaseValue: c.value,
		Value:     c.value,
		Tick:      0,
	}
}

func (c *ConstantSimulator) AddStateSubscriber(subscriber StateSubscriber) {}

func (c *ConstantSimulator) Restart() {}

func (c *ConstantSimulator) ToDTO() dto.Base {
	return dto.Base{
		Type: dto.BaseTypeConstant,
		Data: dto.ConstantBase{
			Value: c.value,
		},
	}
}
