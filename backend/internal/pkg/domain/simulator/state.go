package simulator

type SimulatorState struct {
	PreviousPoint float64
	NextPoint     float64
	TicksDistance uint64

	PointState
}

type PointState struct {
	Value float64
	Tick  uint64
}
