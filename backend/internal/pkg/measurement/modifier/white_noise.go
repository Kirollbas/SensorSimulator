package modifier

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/simulator"
)

type Generator interface {
	NextZeroToOne() float64
}

type WhiteNoise struct {
	prng      Generator
	maxOffset float64
}

func NewWhiteNoiseModifier(
	prng Generator,
	maxOffset float64,
) (*WhiteNoise, error) {
	if prng == nil {
		return nil, fmt.Errorf("generator cannot be nil")
	}

	return &WhiteNoise{
		prng:      prng,
		maxOffset: maxOffset,
	}, nil
}

func (wn *WhiteNoise) ApplyModifier(point simulator.PointState) simulator.PointState {
	point.Value += wn.maxOffset*2*wn.prng.NextZeroToOne() - 1
	return point
}
