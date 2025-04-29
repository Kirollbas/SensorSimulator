package node

import "github.com/gopcua/opcua/server"

type Node struct {
	baseNode     *server.Node
	modifiedNode *server.Node
	namespace    *server.NodeNameSpace
}

func NewNode(namespace *server.NodeNameSpace, baseNode, modifiedNode *server.Node) *Node {
	return &Node{
		namespace:    namespace,
		baseNode:     baseNode,
		modifiedNode: modifiedNode,
	}
}
