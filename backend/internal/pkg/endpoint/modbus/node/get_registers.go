package node

import "sync/atomic"

func (n *SimulatorNode) GetRegisters() (res []uint16) {
	base := atomic.LoadUint64(&n.baseValue)
	modified := atomic.LoadUint64(&n.modifiedValue)
	for i := 0; i < 4; i++ {
		res = append(res, uint16((base&0xffff000000000000)>>48))
		base = base << 16
	}
	for i := 0; i < 4; i++ {
		res = append(res, uint16((modified&0xffff000000000000)>>48))
		modified = modified << 16
	}

	return res
}
