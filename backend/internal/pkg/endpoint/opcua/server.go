package opcua

import (
	"github.com/gopcua/opcua/server"
)

type Server struct {
	server          *server.Server
	commonNamespace *server.NodeNameSpace
}

func NewServer() (*Server, error) {
	server := &Server{}
	server.resetServer()

	return server, nil
}
