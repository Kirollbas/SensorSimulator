package simulator

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/base"
	domain "sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
	"sync"
	"time"
)

type Base interface {
	Iterate() domain.PointState
	AddStateSubscriber(base.StateSubscriber)
	Restart()
	ToDTO() dto.Base
}

type Modifier interface {
	ApplyModifier(point domain.PointState) domain.PointState
	Restart()
	ToDTO() dto.Modifier
}

type EndpointUpdater interface {
	Update(state domain.PointState)
}

type Simulator struct {
	name      string
	address   uint16
	duration  time.Duration
	base      Base
	modifiers []Modifier

	frequency     uint64
	ticker        time.Ticker
	endChan       chan struct{}
	stopTimerChan chan struct{}
	wg            sync.WaitGroup
	mx            sync.Mutex

	isWorking            bool
	currentBaseValueBits uint64
	currentValueBits     uint64

	endPointUpdaters []EndpointUpdater
}

func NewSimulator(
	name string,
	address uint16,
	frequency uint64,
	duration time.Duration,
	base Base,
	modifiers []Modifier,
) (*Simulator, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	if base == nil {
		return nil, fmt.Errorf("base can not be nil")
	}

	if frequency == 0 {
		return nil, fmt.Errorf("frequency must be grater then 0")
	}

	return &Simulator{
		name:          name,
		address:       address,
		duration:      duration,
		base:          base,
		modifiers:     modifiers,
		frequency:     frequency,
		ticker:        *time.NewTicker(time.Second / time.Duration(frequency)),
		endChan:       make(chan struct{}),
		stopTimerChan: make(chan struct{}),
	}, nil
}

func (s *Simulator) ToDTO() dto.Simulator {
	modifiers := []dto.Modifier{}
	for _, modifier := range s.modifiers {
		modifiers = append(modifiers, modifier.ToDTO())
	}

	return dto.Simulator{
		Name:      s.name,
		Address:   int(s.address),
		Base:      s.base.ToDTO(),
		Modifiers: modifiers,
		Duration: dto.Duration{
			Duration: s.duration,
		},
	}
}
