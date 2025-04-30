package modifier

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/state"
)

type Simulator interface {
	GetBaseValue() float64
}

type Dependence struct {
	simulator        Simulator
	coefficient      float64
	dependenceCenter float64
}

func NewDependenceModifier(
	simulator Simulator,
	coefficient float64,
	dependenceCenter float64,
) (*Dependence, error) {
	if simulator == nil {
		return nil, fmt.Errorf("dependent simulator can not be nil")
	}

	return &Dependence{
		simulator:        simulator,
		coefficient:      coefficient,
		dependenceCenter: dependenceCenter,
	}, nil
}

func (d *Dependence) Restart() {}

func (d *Dependence) ApplyModifier(point state.PointState) state.PointState {
	point.Value += d.coefficient * (d.simulator.GetBaseValue() - d.dependenceCenter)
	return point
}
