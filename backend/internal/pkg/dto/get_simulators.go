package dto

type GetSimulators struct {
	Simulators []SimulatorWithStatus `json:"simulators"`
}
