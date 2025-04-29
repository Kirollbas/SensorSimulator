package handler

type startAddress uint64

type Node interface {
	GetRegisters() []uint16
	GetRegistersLength() int
}

type Handler struct {
	inputRegisters map[startAddress]Node
}

func NewHandler() *Handler {
	return &Handler{
		inputRegisters: make(map[startAddress]Node),
	}
}
