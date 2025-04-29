package simulator

import (
	"fmt"
	"math"
	"sync/atomic"
)

func (s *Simulator) tick() {
	point := s.base.Iterate()

	for _, modifier := range s.modifiers {
		point = modifier.ApplyModifier(point)
	}

	atomic.StoreUint64(&s.currentValueBits, math.Float64bits(point.Value))
	atomic.StoreUint64(&s.currentBaseValueBits, math.Float64bits(point.BaseValue))

	for _, updater := range s.endPointUpdaters {
		updater.Update(point)
	}

	fmt.Printf("base %f\n", point.BaseValue) //TODO remove
	fmt.Printf("modified %f\n", point.Value)
}
