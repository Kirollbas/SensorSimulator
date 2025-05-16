package modifier

import (
	"math"
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
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

func (q *Quantization) Restart() {}

func (q *Quantization) ApplyModifier(point state.PointState) state.PointState {
	point.Value = q.quant * math.Round(point.Value/q.quant)
	return point
}

func (q *Quantization) ToDTO() dto.Modifier {
	return dto.Modifier{
		Type: dto.ModifierTypeQuantization,
		Data: dto.QuantizationModifier{
			Quant: q.quant,
		},
	}
}
