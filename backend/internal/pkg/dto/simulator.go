package dto

type Simulator struct {
	Name      string     `json:"name"`
	Address   int        `json:"address"`
	Base      Base       `json:"base"`
	Modifiers []Modifier `json:"modifiers,omitempty"`
	Duration  Duration   `json:"duration,omitempty"`
}

type SimulatorWithStatus struct {
	Simulator
	IsActive bool `json:"is_active"`
}
