package simulator

import "time"

func (s *Simulator) Start() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.isWorking {
		return
	}

	s.wg.Add(1)

	go func() {
		for {
			select {
			case <-s.endChan:
				s.wg.Done()
				return
			case <-s.ticker.C:
				s.tick()
			}
		}
	}()

	timer := time.NewTimer(s.duration)
	go func() {
		for {
			select {
			case <-timer.C:
				if s.duration > 0 {
					s.mx.Lock()
					defer s.mx.Unlock()

					if !s.isWorking {
						return
					}

					s.endChan <- struct{}{}
					s.wg.Wait()

					s.isWorking = false
					return
				}
			case <-s.stopTimerChan:
				return
			}
		}
	}()

	s.isWorking = true
}
