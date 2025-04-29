package modifier

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/simulator"
)

type RandomFixedDash struct {
	prng       Generator
	fixedValue float64

	maxDashTicks   uint64
	minDashTicks   uint64
	avgTicksPeriod uint64

	dashTicksLeft uint64
}

func NewRandomFixedDashModifier(
	prng Generator,
	fixedValue float64,
	maxDashTicks uint64,
	minDashTicks uint64,
	avgTicksPeriod uint64,
) (*RandomFixedDash, error) {
	if prng == nil {
		return nil, fmt.Errorf("generator cannot be nil")
	}

	if maxDashTicks < minDashTicks {
		return nil, fmt.Errorf("max dash ticks must be greater or equal then min dash ticks")
	}

	if maxDashTicks >= avgTicksPeriod {
		return nil, fmt.Errorf("max dash ticks must be lower then dash period in average ticks")
	}

	return &RandomFixedDash{
		prng:           prng,
		fixedValue:     fixedValue,
		maxDashTicks:   maxDashTicks,
		minDashTicks:   minDashTicks,
		avgTicksPeriod: avgTicksPeriod,
		dashTicksLeft:  0,
	}, nil
}

func (r *RandomFixedDash) Restart() {
	r.prng.Restart()
	r.dashTicksLeft = 0
}

func (r *RandomFixedDash) ApplyModifier(point simulator.PointState) simulator.PointState {
	if r.prng.NextZeroToOne() < 1.0/float64(r.avgTicksPeriod) {
		r.dashTicksLeft = uint64(float64(r.maxDashTicks-r.minDashTicks)*r.prng.NextZeroToOne()) + r.minDashTicks
	}

	if r.dashTicksLeft > 0 {
		point.Value = r.fixedValue
		r.dashTicksLeft--
	}

	return point
}
