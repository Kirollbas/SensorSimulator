package base

import (
	"fmt"
	"math"
)

type SineWaveSimulator struct {
	prng Generator

	minValue            float64
	maxValue            float64
	maxTicksUntilChange uint64

	currentValue float64

	startPoint       float64
	endPoint         float64
	distanceTicks    uint64
	ticksUntilChange uint64
}

func NewSineWaveSimulator(
	prng Generator,
	minValue float64,
	maxValue float64,
	maxTicksUntilChange uint64,
) (*SineWaveSimulator, error) {
	if prng == nil {
		return nil, fmt.Errorf("generator cannot be nil")
	}

	if maxTicksUntilChange == 0 {
		return nil, fmt.Errorf("max ticks until next point change must be grater then zero")
	}

	if minValue > maxValue {
		return nil, fmt.Errorf("min value must be lower then max value")
	}

	centerValue := (maxValue - minValue) / 2

	return &SineWaveSimulator{
		prng:                prng,
		minValue:            minValue,
		maxValue:            maxValue,
		maxTicksUntilChange: maxTicksUntilChange,
		currentValue:        centerValue,
		startPoint:          centerValue,
		endPoint:            centerValue,
	}, nil
}

func (s *SineWaveSimulator) Iterate() float64 {
	if s.ticksUntilChange > 0 {
		s.ticksUntilChange--
	} else {
		newDestination := (s.maxValue-s.minValue)*s.prng.NextZeroToOne() + s.minValue
		newTicks := uint(float64(s.maxTicksUntilChange) * s.prng.NextZeroToOne())

		s.distanceTicks = uint64(newTicks)
		s.ticksUntilChange = uint64(newTicks) - 1
		s.startPoint = s.endPoint
		s.endPoint = newDestination
	}

	delta := (s.endPoint - s.startPoint) * math.Sin(
		math.Pi/2*float64(s.distanceTicks-s.ticksUntilChange)/float64(s.distanceTicks),
	)

	newValue := s.startPoint + delta

	s.currentValue = newValue
	return newValue
}
