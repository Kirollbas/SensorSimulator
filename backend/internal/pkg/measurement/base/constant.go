package base

import "sensor-simulator/internal/pkg/domain/simulator"

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

func (c *ConstantSimulator) Iterate() simulator.PointState {
	return simulator.PointState{
		Value: c.value,
		Tick:  0,
	}
}
