package node

import (
	"sensor-simulator/internal/pkg/domain/simulator"
	"time"

	"github.com/gopcua/opcua/ua"
)

func (n *Node) Update(state simulator.PointState) {
	timestamp := time.Now()

	n.baseNode.SetAttribute(ua.AttributeIDValue, &ua.DataValue{
		Value:           ua.MustVariant(state.BaseValue),
		SourceTimestamp: timestamp,
		EncodingMask:    ua.DataValueValue | ua.DataValueSourceTimestamp,
	})

	n.modifiedNode.SetAttribute(ua.AttributeIDValue, &ua.DataValue{
		Value:           ua.MustVariant(state.Value),
		SourceTimestamp: timestamp,
		EncodingMask:    ua.DataValueValue | ua.DataValueSourceTimestamp,
	})

	n.namespace.ChangeNotification(n.baseNode.ID())
	n.namespace.ChangeNotification(n.modifiedNode.ID())
}
