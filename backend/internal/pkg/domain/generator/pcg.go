package generator

import (
	"math"
	"sensor-simulator/internal/pkg/dto"

	"github.com/db47h/rand64/v3/pcg"
)

type Pcg struct {
	seed int64
	gen  pcg.Rng
}

func NewPcg(seed int64) *Pcg {
	pcg := Pcg{
		seed: seed,
		gen:  pcg.Rng{},
	}
	pcg.gen.Seed(seed)

	return &pcg
}

func (x *Pcg) NextInt() int64 {
	return x.gen.Int63()
}

func (x *Pcg) NextFloat() float64 {
	return float64(x.NextInt())
}

func (x *Pcg) NextZeroToOne() float64 {
	return float64(x.NextInt()) / float64(math.MaxInt64)
}

func (x *Pcg) Restart() {
	x.gen.Seed(x.seed)
}

func (x *Pcg) ToDTO() dto.Prng {
	return dto.Prng{
		Type: dto.PRNGTypePCG,
		Seed: x.seed,
	}
}
