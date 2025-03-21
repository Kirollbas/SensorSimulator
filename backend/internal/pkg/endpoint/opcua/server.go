package opcua

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/server"
	"github.com/gopcua/opcua/server/attrs"
	"github.com/gopcua/opcua/ua"
)

func Start() {

	var opts []server.Option
	opts = append(opts,
		server.EndPoint("localhost", 4840),
		server.EnableSecurity("None", ua.MessageSecurityModeNone),
		server.EnableAuthMode(ua.UserTokenTypeAnonymous),
	)

	s := server.New(
		opts...,
	)

	// add the namespaces to the server, and add a reference to them if desired.
	// here we are choosing to add the namespaces to the root/object folder
	// to do this we first need to get the root namespace object folder so we
	// get the object node
	root_ns, _ := s.Namespace(0)
	root_obj_node := root_ns.Objects()

	// Start the server
	// Note that you can add namespaces before or after starting the server.
	if err := s.Start(context.Background()); err != nil {
		fmt.Printf("Error starting server, exiting: %s", err)
	}
	defer s.Close()
	s.Session(&ua.RequestHeader{})

	// Now we'll add a node namespace.  This is a more traditional way to add nodes to the server
	// and is more in line with the opc ua node model, but may be more cumbersome for some use cases.
	nodeNS := server.NewNodeNameSpace(s, "NodeNamespace")
	fmt.Printf("Node Namespace added at index %d", nodeNS.ID())

	// add the reference for this namespace's root object folder to the server's root object folder
	// but you can add a reference to whatever node(s) you need
	nns_obj := nodeNS.Objects()
	root_obj_node.AddRef(nns_obj, id.HasComponent, true)

	// Create some nodes for it.  Here we are usin gthe AddNewVariableNode utility function to create a new variable node
	// with an integer node ID that is automatically assigned. (ns=<namespace id>,s=<auto assigned>)
	// be sure to add the reference to the node somewhere if desired, or clients won't be able to browse it.
	var1 := nodeNS.AddNewVariableNode("TestVar1", float32(123.45))
	nns_obj.AddRef(var1, id.HasComponent, true)

	// This node will have a string node id (ns=<namespace id>,s=TestVar2)
	// your variable node's value can also return a ua.Variant from a function if you want to update the value dynamically
	// here we are just incrementing a counter every time the value is read.

	// Now we'll add a node from scratch.  This is a more manual way to add nodes to the server and gives you full
	// control, but you'll have to build the node up with the correct attributes and references and then reference it from
	// the parent node in the namespace if applicable.
	var3 := server.NewNode(
		ua.NewNumericNodeID(nodeNS.ID(), 12345), // you can use whatever node id you want here, whether it's numeric, string, guid, etc...
		map[ua.AttributeID]*ua.DataValue{
			ua.AttributeIDBrowseName: server.DataValueFromValue(attrs.BrowseName("MyBrowseName")),
			ua.AttributeIDNodeClass:  server.DataValueFromValue(uint32(ua.NodeClassVariable)),
		},
		nil,
		func() *ua.DataValue { return server.DataValueFromValue(12.34) },
	)
	nodeNS.AddNode(var3)
	nns_obj.AddRef(var3, id.HasComponent, true)

	// simulate a background process updating the data in the namespace.
	go func() {
		updates := 0
		num := 42
		time.Sleep(time.Second * 10)
		for {
			updates++
			num++

			// get the current value of the variable
			last_value := var1.Value().Value.Value().(float32)
			// and change it
			last_value += 1

			// wrap the new value in a DataValue and use that to update the Value attribute of the node
			val := ua.DataValue{
				Value:           ua.MustVariant(last_value),
				SourceTimestamp: time.Now(),
				EncodingMask:    ua.DataValueValue | ua.DataValueSourceTimestamp,
			}
			var1.SetAttribute(ua.AttributeIDValue, &val)

			// we also need to let the node namespace know that the value has changed so it can trigger the change notification
			// and send the updated value to any subscribed clients.
			nodeNS.ChangeNotification(var1.ID())

			time.Sleep(time.Second)
		}
	}()

	// simulate monitoring one of the namespaces for change events.
	// this is how you would be notified when a write to a node
	// occurs through the opc ua server
	go func() {
		for {
			changed_id := <-nodeNS.ExternalNotification
			node := nodeNS.Node(changed_id)
			value := node.Value().Value.Value()
			fmt.Printf("%s changed to %v", changed_id.String(), value)
		}
	}()

	// catch ctrl-c and gracefully shutdown the server.
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	defer signal.Stop(sigch)
	fmt.Printf("Press CTRL-C to exit")

	<-sigch
	fmt.Printf("Shutting down the server...")
}
