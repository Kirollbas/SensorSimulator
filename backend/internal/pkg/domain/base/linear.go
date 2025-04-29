package base

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/state"
)

type LinearSimulator struct {
	prng             Generator
	stateSubscribers []StateSubscriber

	minValue            float64
	maxValue            float64
	minTicksUntilChange uint64
	maxTicksUntilChange uint64

	currentValue float64

	nextPoint     float64
	speed         float64
	distanceTicks uint64
	currentTick   uint64
}

func NewLinearSimulator(
	prng Generator,
	minValue float64,
	maxValue float64,
	minTicksUntilChange uint64,
	maxTicksUntilChange uint64,
) (*LinearSimulator, error) {
	if prng == nil {
		return nil, fmt.Errorf("generator cannot be nil")
	}

	if maxTicksUntilChange < minTicksUntilChange {
		return nil, fmt.Errorf("min ticks must be lower then max ticks")
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
		stateSubscribers:    make([]StateSubscriber, 0),
		minValue:            minValue,
		maxValue:            maxValue,
		minTicksUntilChange: minTicksUntilChange,
		maxTicksUntilChange: maxTicksUntilChange,
		currentValue:        centerValue,
		nextPoint:           centerValue,
	}, nil
}

func (l *LinearSimulator) Restart() {
	centerValue := (l.maxValue - l.minValue) / 2

	l.currentValue = centerValue
	l.nextPoint = centerValue
	l.speed = 0
	l.distanceTicks = 0
	l.currentTick = 0
}

func (l *LinearSimulator) AddStateSubscriber(subscriber StateSubscriber) {
	l.stateSubscribers = append(l.stateSubscribers, subscriber)
}

func (l *LinearSimulator) Iterate() state.PointState {
	if l.currentTick >= l.distanceTicks {
		newDestination := (l.maxValue-l.minValue)*l.prng.NextZeroToOne() + l.minValue
		newTicks := uint64(float64(l.maxTicksUntilChange-l.minTicksUntilChange)*l.prng.NextZeroToOne()) + l.minTicksUntilChange

		state := state.SimulatorBaseState{
			PreviousPoint: l.nextPoint,
			NextPoint:     newDestination,
			TicksDistance: uint64(newTicks),
		}
		for _, subscriber := range l.stateSubscribers {
			subscriber.UpdateState(state)
		}

		l.distanceTicks = uint64(newTicks)
		l.currentTick = 0

		distance := newDestination - l.nextPoint
		l.speed = distance / float64(newTicks)
		l.currentValue = l.nextPoint
		l.nextPoint = newDestination
	}

	l.currentValue += l.speed
	l.currentTick++

	return state.PointState{
		BaseValue: l.currentValue,
		Value:     l.currentValue,
		Tick:      l.currentTick,
	}
}
