package opcua

import (
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/ua"
)

type Server struct {
	server          *server.Server
	commonNamespace *server.NodeNameSpace
}

func NewServer() (*Server, error) {
	host := "localhost"
	port := 46010

	var opts []server.Option
	opts = append(opts,
		server.EndPoint(host, port),
		server.EnableSecurity("None", ua.MessageSecurityModeNone),
		server.EnableAuthMode(ua.UserTokenTypeAnonymous),
		server.SoftwareVersion("0.0.1"),
	)

	opcuaServer := server.New(opts...)

	nodeNS := server.NewNodeNameSpace(opcuaServer, "SensorSimulators")
	nns_obj := nodeNS.Objects()

	root_ns, _ := opcuaServer.Namespace(0)
	root_obj_node := root_ns.Objects()
	root_obj_node.AddRef(nns_obj, id.HasComponent, true)

	return &Server{
		server:          opcuaServer,
		commonNamespace: nodeNS,
	}, nil
}
