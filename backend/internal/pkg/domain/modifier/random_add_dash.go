package modifier

import (
	"fmt"
	"sensor-simulator/internal/pkg/domain/state"
	"sensor-simulator/internal/pkg/dto"
	"time"
)

type RandomAddDash struct {
	prng            Generator
	minAddValue     float64
	maxAddValue     float64
	maxDashDuration time.Duration
	minDashDuration time.Duration
	avgPeriod       time.Duration

	maxDashTicks   uint64
	minDashTicks   uint64
	avgTicksPeriod uint64

	addValue      float64
	dashTicksLeft uint64
}

func NewRandomAddDashModifier(
	prng Generator,
	maxDashDuration time.Duration,
	minDashDuration time.Duration,
	avgPeriod time.Duration,
	minAddValue float64,
	maxAddValue float64,
	tickPeriod time.Duration,
) (*RandomAddDash, error) {
	minDashTicks := uint64(minDashDuration / tickPeriod)
	maxDashTicks := uint64(maxDashDuration / tickPeriod)
	avgDashTicks := uint64(avgPeriod / tickPeriod)

	if prng == nil {
		return nil, fmt.Errorf("generator cannot be nil")
	}

	if maxAddValue < minAddValue {
		return nil, fmt.Errorf("max add value must be greater or equal then min add value")
	}

	if maxDashTicks < minDashTicks {
		return nil, fmt.Errorf("max dash ticks must be greater or equal then min dash ticks")
	}

	if maxDashTicks >= avgDashTicks {
		return nil, fmt.Errorf("max dash ticks must be lower then dash period in average ticks")
	}

	return &RandomAddDash{
		prng:            prng,
		maxDashTicks:    maxDashTicks,
		minDashTicks:    minDashTicks,
		maxDashDuration: maxDashDuration,
		minDashDuration: minDashDuration,
		avgPeriod:       avgPeriod,
		avgTicksPeriod:  avgDashTicks,
		minAddValue:     minAddValue,
		maxAddValue:     maxAddValue,
	}, nil
}

func (r *RandomAddDash) Restart() {
	r.prng.Restart()
	r.addValue = 0
	r.dashTicksLeft = 0
}

func (r *RandomAddDash) ApplyModifier(point state.PointState) state.PointState {
	if r.prng.NextZeroToOne() < 1.0/float64(r.avgTicksPeriod) {
		r.dashTicksLeft = uint64(float64(r.maxDashTicks-r.minDashTicks)*r.prng.NextZeroToOne()) + r.minDashTicks
		r.addValue = (r.maxAddValue-r.minAddValue)*r.prng.NextZeroToOne() + r.minAddValue
	}

	if r.dashTicksLeft > 0 {
		point.Value += r.addValue
		r.dashTicksLeft--
	}

	return point
}

func (r *RandomAddDash) ToDTO() dto.Modifier {
	return dto.Modifier{
		Type: dto.ModifierTypeRandomAddDash,
		Data: dto.RandomAddDashModifier{
			Generator:   r.prng.ToDTO(),
			MinAddValue: r.minAddValue,
			MaxAddValue: r.maxAddValue,
			MinDashDuration: dto.Duration{
				Duration: r.minDashDuration,
			},
			MaxDashDuration: dto.Duration{
				Duration: r.maxDashDuration,
			},
			AvgPeriod: dto.Duration{
				Duration: r.avgPeriod,
			},
		},
	}
}
