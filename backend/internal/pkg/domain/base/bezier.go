package base

import (
	"fmt"
	"math"
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
	"time"
)

type Generator interface {
	NextZeroToOne() float64
	Restart()
	ToDTO() dto.Prng
}

type StateSubscriber interface {
	UpdateState(state state.SimulatorBaseState)
}

type BezierSimulator struct {
	prng             Generator
	stateSubscribers []StateSubscriber

	minValue  float64
	maxValue  float64
	minPeriod time.Duration
	maxPeriod time.Duration

	minTicksUntilChange uint64
	maxTicksUntilChange uint64

	currentValue float64

	startPoint    float64
	endPoint      float64
	distanceTicks uint64
	currentTick   uint64
}

func NewBezierSimulator(
	prng Generator,
	minValue float64,
	maxValue float64,
	minPeriod time.Duration,
	maxPeriod time.Duration,
	tickPeriod time.Duration,
) (*BezierSimulator, error) {
	minTicksUntilChange := uint64(minPeriod / tickPeriod)
	maxTicksUntilChange := uint64(maxPeriod / tickPeriod)

	if prng == nil {
		return nil, fmt.Errorf("generator cannot be nil")
	}

	if maxTicksUntilChange == 0 {
		return nil, fmt.Errorf("max ticks until next point change must be grater then zero")
	}

	if maxTicksUntilChange < minTicksUntilChange {
		return nil, fmt.Errorf("min ticks must be lower then max ticks")
	}

	if minValue > maxValue {
		return nil, fmt.Errorf("min value must be lower then max value")
	}

	centerValue := (maxValue - minValue) / 2

	return &BezierSimulator{
		prng:                prng,
		stateSubscribers:    make([]StateSubscriber, 0),
		minValue:            minValue,
		maxValue:            maxValue,
		minPeriod:           minPeriod,
		maxPeriod:           maxPeriod,
		minTicksUntilChange: minTicksUntilChange,
		maxTicksUntilChange: maxTicksUntilChange,
		currentValue:        centerValue,
		startPoint:          centerValue,
		endPoint:            centerValue,
	}, nil
}

func (b *BezierSimulator) Restart() {
	centerValue := (b.maxValue - b.minValue) / 2

	b.prng.Restart()
	b.currentTick = 0
	b.distanceTicks = 0
	b.currentValue = centerValue
	b.startPoint = centerValue
	b.endPoint = centerValue
}

func (b *BezierSimulator) AddStateSubscriber(subscriber StateSubscriber) {
	b.stateSubscribers = append(b.stateSubscribers, subscriber)
}

func (b *BezierSimulator) Iterate() state.PointState {
	if b.currentTick >= b.distanceTicks {
		newDestination := (b.maxValue-b.minValue)*b.prng.NextZeroToOne() + b.minValue
		newTicks := uint64(float64(b.maxTicksUntilChange-b.minTicksUntilChange)*b.prng.NextZeroToOne()) + b.minTicksUntilChange

		state := state.SimulatorBaseState{
			PreviousPoint: b.endPoint,
			NextPoint:     newDestination,
			TicksDistance: newTicks,
		}
		for _, subscriber := range b.stateSubscribers {
			subscriber.UpdateState(state)
		}

		b.distanceTicks = uint64(newTicks)
		b.currentTick = 0
		b.startPoint = b.endPoint
		b.endPoint = newDestination
	}

	currentTime := float64(b.currentTick) / float64(b.distanceTicks)

	newValue := math.Pow(1-currentTime, 3)*b.startPoint +
		3*math.Pow(1-currentTime, 2)*(currentTime)*b.startPoint +
		3*(1-currentTime)*math.Pow(currentTime, 2)*b.endPoint +
		math.Pow(currentTime, 3)*b.endPoint

	b.currentValue = newValue
	b.currentTick++

	return state.PointState{
		BaseValue: newValue,
		Value:     newValue,
		Tick:      b.currentTick,
	}
}

func (b *BezierSimulator) ToDTO() dto.Base {
	return dto.Base{
		Type: dto.BaseTypeBezier,
		Data: dto.CommonBase{
			Generator: b.prng.ToDTO(),
			MinValue:  b.minValue,
			MaxValue:  b.maxValue,
			MinPeriod: dto.Duration{
				Duration: b.minPeriod,
			},
			MaxPeriod: dto.Duration{
				Duration: b.maxPeriod,
			},
		},
	}
}
