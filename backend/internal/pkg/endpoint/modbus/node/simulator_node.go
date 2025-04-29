package node

type SimulatorNode struct {
	baseValue     uint64
	modifiedValue uint64
}

func NewSimulatorNode() *SimulatorNode {
	return &SimulatorNode{}
}

func (n SimulatorNode) GetRegistersLength() int {
	return 8
}
