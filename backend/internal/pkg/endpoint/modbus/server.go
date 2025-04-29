package modbus

import (
	"fmt"
	"sensor-simulator/internal/pkg/endpoint/modbus/handler"
	"time"

	"github.com/simonvetter/modbus"
)

type Server struct {
	server  *modbus.ModbusServer
	handler *handler.Handler
}

func NewServer() (*Server, error) {
	handler := handler.NewHandler()

	server, err := modbus.NewServer(&modbus.ServerConfiguration{
		URL:        "tcp://localhost:5502",
		Timeout:    30 * time.Second,
		MaxClients: 5,
	}, handler)
	if err != nil {
		return nil, fmt.Errorf("failed to create modbus server. Err: %w", err)
	}

	return &Server{
		server:  server,
		handler: handler,
	}, nil
}
