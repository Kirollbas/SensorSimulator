package state

type SimulatorBaseState struct {
	PreviousPoint float64
	NextPoint     float64
	TicksDistance uint64
}

type PointState struct {
	BaseValue float64
	Value     float64
	Tick      uint64
}
