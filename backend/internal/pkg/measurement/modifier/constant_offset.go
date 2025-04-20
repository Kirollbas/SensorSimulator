package modifier

type ConstantOffset struct {
	value float64
}

func NewOffsetModifier(
	offset float64,
) (*ConstantOffset, error) {
	return &ConstantOffset{
		value: offset,
	}, nil
}

func (o *ConstantOffset) ApplyModifier(value float64) float64 {
	return value + o.value
}
