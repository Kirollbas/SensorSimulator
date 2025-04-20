package base

import (
	"fmt"
)

type LinearSimulator struct {
	prng Generator

	minValue            float64
	maxValue            float64
	maxTicksUntilChange uint64

	currentValue float64

	nextPoint        float64
	speed            float64
	distanceTicks    uint64
	ticksUntilChange uint64
}

func NewLinearSimulator(
	prng Generator,
	minValue float64,
	maxValue float64,
	maxTicksUntilChange uint64,
) (*LinearSimulator, error) {
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

	return &LinearSimulator{
		prng:                prng,
		minValue:            minValue,
		maxValue:            maxValue,
		maxTicksUntilChange: maxTicksUntilChange,
		currentValue:        centerValue,
		nextPoint:           centerValue,
	}, nil
}

func (l *LinearSimulator) Iterate() float64 {
	if l.ticksUntilChange > 0 {
		l.ticksUntilChange--
	} else {
		newDestination := (l.maxValue-l.minValue)*l.prng.NextZeroToOne() + l.minValue
		newTicks := uint(float64(l.maxTicksUntilChange) * l.prng.NextZeroToOne())

		l.distanceTicks = uint64(newTicks)
		l.ticksUntilChange = uint64(newTicks) - 1

		distance := newDestination - l.nextPoint
		l.speed = distance / float64(newTicks)
		l.currentValue = l.nextPoint
		l.nextPoint = newDestination
	}

	l.currentValue += l.speed
	return l.currentValue
}
