package modifier

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
	"time"
)

type RandomFixedDash struct {
	prng            Generator
	fixedValue      float64
	maxDashDuration time.Duration
	minDashDuration time.Duration
	avgPeriod       time.Duration

	maxDashTicks   uint64
	minDashTicks   uint64
	avgTicksPeriod uint64

	dashTicksLeft uint64
}

func NewRandomFixedDashModifier(
	prng Generator,
	fixedValue float64,
	maxDashDuration time.Duration,
	minDashDuration time.Duration,
	avgPeriod time.Duration,
	tickPeriod time.Duration,
) (*RandomFixedDash, error) {
	minDashTicks := uint64(minDashDuration / tickPeriod)
	maxDashTicks := uint64(maxDashDuration / tickPeriod)
	avgDashTicks := uint64(avgPeriod / tickPeriod)

	if prng == nil {
		return nil, fmt.Errorf("generator cannot be nil")
	}

	if maxDashTicks < minDashTicks {
		return nil, fmt.Errorf("max dash ticks must be greater or equal then min dash ticks")
	}

	if maxDashTicks >= avgDashTicks {
		return nil, fmt.Errorf("max dash ticks must be lower then dash period in average ticks")
	}

	return &RandomFixedDash{
		prng:            prng,
		fixedValue:      fixedValue,
		maxDashDuration: maxDashDuration,
		minDashDuration: minDashDuration,
		avgPeriod:       avgPeriod,
		maxDashTicks:    maxDashTicks,
		minDashTicks:    minDashTicks,
		avgTicksPeriod:  avgDashTicks,
		dashTicksLeft:   0,
	}, nil
}

func (r *RandomFixedDash) Restart() {
	r.prng.Restart()
	r.dashTicksLeft = 0
}

func (r *RandomFixedDash) ApplyModifier(point state.PointState) state.PointState {
	if r.prng.NextZeroToOne() < 1.0/float64(r.avgTicksPeriod) {
		r.dashTicksLeft = uint64(float64(r.maxDashTicks-r.minDashTicks)*r.prng.NextZeroToOne()) + r.minDashTicks
	}

	if r.dashTicksLeft > 0 {
		point.Value = r.fixedValue
		r.dashTicksLeft--
	}

	return point
}

func (r *RandomFixedDash) ToDTO() dto.Modifier {
	return dto.Modifier{
		Type: dto.ModifierTypeRandomFixedDash,
		Data: dto.RandomFixedDashModifier{
			Generator: r.prng.ToDTO(),
			MinDashDuration: dto.Duration{
				Duration: r.minDashDuration,
			},
			MaxDashDuration: dto.Duration{
				Duration: r.maxDashDuration,
			},
			AvgPeriod: dto.Duration{
				Duration: r.avgPeriod,
			},
			Value: r.fixedValue,
		},
	}
}
