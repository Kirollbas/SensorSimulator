package handler

func (h *Handler) ClearNodes() {
	h.inputRegisters = make(map[startAddress]Node)
}
