package handler

import "github.com/simonvetter/modbus"

func (h *Handler) HandleCoils(req *modbus.CoilsRequest) (res []bool, err error) {
	return nil, modbus.ErrIllegalFunction
}
