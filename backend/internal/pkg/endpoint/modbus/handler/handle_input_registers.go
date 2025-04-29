package handler

import (
	"math"
	"sensor-simulator/internal/pkg/utils"

	"github.com/simonvetter/modbus"
)

func (h *Handler) HandleInputRegisters(req *modbus.InputRegistersRequest) (res []uint16, err error) {
	if req.Quantity == 0 {
		return
	}

	if math.MaxUint16-req.Addr < req.Quantity-1 {
		err = modbus.ErrIllegalDataAddress
		return
	}

	for i := 0; i < int(req.Quantity); i++ {
		address := int(req.Addr) + i

		if _, ok := h.inputRegisters[startAddress(address)]; !ok {
			res = append(res, 0x0000)
			continue
		}

		node := h.inputRegisters[startAddress(address)]
		registersCount := utils.MinInt(node.GetRegistersLength(), int(req.Quantity)-i)
		res = append(res, node.GetRegisters()[0:registersCount]...)
		i += registersCount - 1
	}

	return
}
