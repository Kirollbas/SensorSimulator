package modifier

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
)

type Simulator interface {
	GetBaseValue() float64
	GetName() string
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

func (d *Dependence) ToDTO() dto.Modifier {
	return dto.Modifier{
		Type: dto.ModifierTypeDependence,
		Data: dto.DependenceModifier{
			SimulatorName: d.simulator.GetName(),
			Center:        d.dependenceCenter,
			Coefficient:   d.coefficient,
		},
	}
}
