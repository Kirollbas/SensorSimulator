package simulator

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/base"
	domain "sensor-simulator/internal/pkg/domain/state"
	"sync"
	"time"
)

type Base interface {
	Iterate() domain.PointState
	AddStateSubscriber(base.StateSubscriber)
	Restart()
}

type Modifier interface {
	ApplyModifier(point domain.PointState) domain.PointState
	Restart()
}

type EndpointUpdater interface {
	Update(state domain.PointState)
}

type Simulator struct {
	name      string
	address   uint16
	base      Base
	modifiers []Modifier

	frequency uint64
	ticker    time.Ticker
	endChan   chan struct{}
	wg        sync.WaitGroup

	currentBaseValueBits uint64
	currentValueBits     uint64

	endPointUpdaters []EndpointUpdater
}

func NewSimulator(
	name string,
	address uint16,
	frequency uint64,
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
		name:      name,
		address:   address,
		base:      base,
		modifiers: modifiers,
		frequency: frequency,
		ticker:    *time.NewTicker(time.Second / time.Duration(frequency)),
		endChan:   make(chan struct{}),
	}, nil
}
