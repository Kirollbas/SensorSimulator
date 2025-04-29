package simulator

func (s *Simulator) Restart() {
	s.wg.Wait()

	s.base.Restart()
	for _, modifier := range s.modifiers {
		modifier.Restart()
	}

	s.endPointUpdaters = []EndpointUpdater{}
}
