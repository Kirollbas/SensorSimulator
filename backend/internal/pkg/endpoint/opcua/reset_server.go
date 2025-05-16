package opcua

import (
	"sensor-simulator/internal/configs"

	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/ua"
)

func (s *Server) resetServer() {
	host := configs.GetConfig().Opcua.Host
	port := configs.GetConfig().Opcua.Port
	namespaceName := configs.GetConfig().Opcua.SimulatorNamespaceName

	var opts []server.Option
	opts = append(opts,
		server.EndPoint(host, port),
		server.EnableSecurity("None", ua.MessageSecurityModeNone),
		server.EnableAuthMode(ua.UserTokenTypeAnonymous),
		server.SoftwareVersion("0.0.1"),
	)

	opcuaServer := server.New(opts...)

	nodeNS := server.NewNodeNameSpace(opcuaServer, namespaceName)
	nns_obj := nodeNS.Objects()

	root_ns, _ := opcuaServer.Namespace(0)
	root_obj_node := root_ns.Objects()
	root_obj_node.AddRef(nns_obj, id.HasComponent, true)

	s.server = opcuaServer
	s.commonNamespace = nodeNS
}
