package modifier

import (
	"math"
	"sensor-simulator/internal/pkg/domain/simulator"
)

type NonLinear struct {
	coefficient float64
	center      float64
}

func NewNonLinearModifier(
	coefficient float64,
	center float64,
) (*NonLinear, error) {
	return &NonLinear{
		coefficient: coefficient,
		center:      center,
	}, nil
}

func (nl *NonLinear) Restart() {}

func (nl *NonLinear) ApplyModifier(point simulator.PointState) simulator.PointState {
	diff := point.Value - nl.center

	point.Value += nl.coefficient * math.Pow(diff, 2) * math.Copysign(1.0, diff)

	return point
}
