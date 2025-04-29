package simulator

func (s *Simulator) Finish() {
	s.endChan <- struct{}{}
	s.wg.Wait()
}
