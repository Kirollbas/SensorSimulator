package simulator

import (
	"fmt"
	"math"
	"sensor-simulator/internal/configs"
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

	if configs.GetConfig().Simulator.LogsEnabled {
		s.logFile.WriteString(fmt.Sprintf("base %f\n", point.BaseValue))
		s.logFile.WriteString(fmt.Sprintf("modified %f\n", point.Value))
	}
}
