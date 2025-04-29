package handler

import "github.com/simonvetter/modbus"

func (h *Handler) HandleHoldingRegisters(req *modbus.HoldingRegistersRequest) (res []uint16, err error) {
	return nil, modbus.ErrIllegalFunction
}
