package simulator

import "sensor-simulator/internal/configs"

func (s *Simulator) Finish() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if !s.isWorking {
		return
	}

	if configs.GetConfig().Simulator.LogsEnabled {
		s.logFile.Close()
	}

	s.endChan <- struct{}{}
	s.wg.Wait()

	s.stopTimerChan <- struct{}{}
	s.isWorking = false
}
