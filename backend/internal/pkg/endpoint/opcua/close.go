package opcua

import (
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/ua"
)

func (s *Server) Close() {
	s.server.Close()

	host := "localhost"
	port := 46010

	var opts []server.Option
	opts = append(opts,
		server.EndPoint(host, port),
		server.EnableSecurity("None", ua.MessageSecurityModeNone),
		server.EnableAuthMode(ua.UserTokenTypeAnonymous),
		server.SoftwareVersion("0.0.1"),
	)

	s.server = server.New(opts...)

	nodeNS := server.NewNodeNameSpace(s.server, "SensorSimulators")
	nns_obj := nodeNS.Objects()

	root_ns, _ := s.server.Namespace(0)
	root_obj_node := root_ns.Objects()
	root_obj_node.AddRef(nns_obj, id.HasComponent, true)

	s.commonNamespace = nodeNS
}
