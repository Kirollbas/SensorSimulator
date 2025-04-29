package simulator

func (s *Simulator) AddEndpointUpdater(updater EndpointUpdater) {
	s.endPointUpdaters = append(s.endPointUpdaters, updater)
}
