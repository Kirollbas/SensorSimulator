package base

import (
	"sensor-simulator/internal/pkg/domain/state"
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
