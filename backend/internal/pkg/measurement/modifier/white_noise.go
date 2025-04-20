package modifier

import "fmt"

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

func (wn *WhiteNoise) ApplyModifier(value float64) float64 {
	return value + wn.maxOffset*2*wn.prng.NextZeroToOne() - 1
}
