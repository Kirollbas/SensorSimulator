package generator

import (
	"math"

	"github.com/db47h/rand64/v3/pcg"
)

type Pcg struct {
	gen pcg.Rng
}

func NewPcg(seed int64) *Pcg {
	pcg := Pcg{
		gen: pcg.Rng{},
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
