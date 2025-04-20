package generator

import (
	"math"

	"github.com/db47h/rand64/v3/xoshiro"
)

type Xoshiro struct {
	gen xoshiro.Rng256P
}

func NewXoshiro(seed int64) *Xoshiro {
	xoshiro := Xoshiro{
		gen: xoshiro.Rng256P{},
	}
	xoshiro.gen.Seed(seed)

	return &xoshiro
}

func (x *Xoshiro) NextInt() int64 {
	return x.gen.Int63()
}

func (x *Xoshiro) NextFloat() float64 {
	return float64(x.NextInt())
}

func (x *Xoshiro) NextZeroToOne() float64 {
	return float64(x.NextInt()) / float64(math.MaxInt64)
}
