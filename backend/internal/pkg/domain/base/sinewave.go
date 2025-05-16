package base

import (
	"fmt"
	"math"
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
	"time"
)

type SineWaveSimulator struct {
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

	currentTick uint64
}

func NewSineWaveSimulator(
	prng Generator,
	minValue float64,
	maxValue float64,
	minPeriod time.Duration,
	maxPeriod time.Duration,
	tickPeriod time.Duration,
) (*SineWaveSimulator, error) {
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

	return &SineWaveSimulator{
		prng:                prng,
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

func (s *SineWaveSimulator) Restart() {
	centerValue := (s.maxValue - s.minValue) / 2

	s.currentValue = centerValue
	s.startPoint = centerValue
	s.endPoint = centerValue
	s.distanceTicks = 0
	s.currentTick = 0
}

func (s *SineWaveSimulator) AddStateSubscriber(subscriber StateSubscriber) {
	s.stateSubscribers = append(s.stateSubscribers, subscriber)
}

func (s *SineWaveSimulator) Iterate() state.PointState {
	if s.currentTick >= s.distanceTicks {
		newDestination := (s.maxValue-s.minValue)*s.prng.NextZeroToOne() + s.minValue
		newTicks := uint64(float64(s.maxTicksUntilChange-s.minTicksUntilChange)*s.prng.NextZeroToOne()) + s.minTicksUntilChange

		state := state.SimulatorBaseState{
			PreviousPoint: s.endPoint,
			NextPoint:     newDestination,
			TicksDistance: uint64(newTicks),
		}
		for _, subscriber := range s.stateSubscribers {
			subscriber.UpdateState(state)
		}

		s.distanceTicks = uint64(newTicks)
		s.currentTick = 0
		s.startPoint = s.endPoint
		s.endPoint = newDestination
	}

	delta := (s.endPoint - s.startPoint) * math.Sin(
		math.Pi/2*float64(s.currentTick)/float64(s.distanceTicks),
	)

	newValue := s.startPoint + delta
	s.currentValue = newValue
	s.currentTick++

	return state.PointState{
		BaseValue: newValue,
		Value:     newValue,
		Tick:      s.currentTick,
	}
}

func (s *SineWaveSimulator) ToDTO() dto.Base {
	return dto.Base{
		Type: dto.BaseTypeSinewave,
		Data: dto.CommonBase{
			Generator: s.prng.ToDTO(),
			MinValue:  s.minValue,
			MaxValue:  s.maxValue,
			MinPeriod: dto.Duration{
				Duration: s.minPeriod,
			},
			MaxPeriod: dto.Duration{
				Duration: s.maxPeriod,
			},
		},
	}
}
