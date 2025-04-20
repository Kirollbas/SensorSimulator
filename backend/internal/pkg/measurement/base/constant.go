package base

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

func (c *ConstantSimulator) Iterate() float64 {
	return c.value
}
