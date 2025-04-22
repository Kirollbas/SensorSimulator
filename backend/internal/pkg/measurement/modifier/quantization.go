package modifier

import (
	"math"
	"sensor-simulator/internal/pkg/domain/simulator"
)

type Quantization struct {
	quant float64
}

func NewQuantizationModifier(
	quant float64,
) (*Quantization, error) {
	return &Quantization{
		quant: quant,
	}, nil
}

func (q *Quantization) ApplyModifier(point simulator.PointState) simulator.PointState {
	point.Value = q.quant * math.Round(point.Value/q.quant)
	return point
}
