package simulator

func (s *Simulator) Start() {
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
}
