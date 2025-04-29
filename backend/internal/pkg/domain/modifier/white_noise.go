package modifier

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/state"
)

type Generator interface {
	NextZeroToOne() float64
	Restart()
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

func (wn *WhiteNoise) Restart() {
	wn.prng.Restart()
}

func (wn *WhiteNoise) ApplyModifier(point state.PointState) state.PointState {
	point.Value += wn.maxOffset*2*wn.prng.NextZeroToOne() - 1
	return point
}
