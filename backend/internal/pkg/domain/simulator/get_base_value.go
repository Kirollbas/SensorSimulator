package simulator

import (
	"math"
	"sync/atomic"
)

func (s *Simulator) GetBaseValue() float64 {
	return math.Float64frombits(atomic.LoadUint64(&s.currentBaseValueBits))
}
