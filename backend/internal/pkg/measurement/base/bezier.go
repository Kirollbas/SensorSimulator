package base

import (
	"fmt"
	"math"
)

type Generator interface {
	NextFloat() float64
	NextZeroToOne() float64
}

type BezierSimulator struct {
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

func NewBezierSimulator(
	prng Generator,
	minValue float64,
	maxValue float64,
	maxTicksUntilChange uint64,
) (*BezierSimulator, error) {
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

	return &BezierSimulator{
		prng:                prng,
		minValue:            minValue,
		maxValue:            maxValue,
		maxTicksUntilChange: maxTicksUntilChange,
		currentValue:        centerValue,
		startPoint:          centerValue,
		endPoint:            centerValue,
	}, nil
}

func (b *BezierSimulator) Iterate() float64 {
	if b.ticksUntilChange > 0 {
		b.ticksUntilChange--
	} else {
		newDestination := (b.maxValue-b.minValue)*b.prng.NextZeroToOne() + b.minValue
		newTicks := uint(float64(b.maxTicksUntilChange) * b.prng.NextZeroToOne())

		b.distanceTicks = uint64(newTicks)
		b.ticksUntilChange = uint64(newTicks) - 1
		b.startPoint = b.endPoint
		b.endPoint = newDestination
	}

	currentTime := float64(b.ticksUntilChange) / float64(b.distanceTicks)

	newValue := math.Pow(currentTime, 3)*b.startPoint +
		3*math.Pow(currentTime, 2)*(1-currentTime)*b.startPoint +
		3*currentTime*math.Pow(1-currentTime, 2)*b.endPoint +
		math.Pow(1-currentTime, 3)*b.endPoint

	b.currentValue = newValue
	return newValue
}
