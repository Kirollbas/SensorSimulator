package handler

import (
	"fmt"
	"math"
)

func (h *Handler) AddNode(address uint16, node Node) error {
	if int(address)+node.GetRegistersLength() > math.MaxUint16 {
		return fmt.Errorf("sum of address and length is over avaliable address")
	}

	for i := int(address); i < int(address)+node.GetRegistersLength(); i++ {
		if _, ok := h.inputRegisters[startAddress(i)]; ok {
			return fmt.Errorf("registers already in use by other node")
		}
	}

	h.inputRegisters[startAddress(address)] = node
	return nil
}
