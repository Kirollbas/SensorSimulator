package modbus

import (
	"fmt"
	"sensor-simulator/internal/config"
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

	config := config.GetConfig()
	url := fmt.Sprintf("tcp://%s:%d", config.Modbus.Host, config.Modbus.Port)

	server, err := modbus.NewServer(&modbus.ServerConfiguration{
		URL:        url,
		Timeout:    time.Duration(config.Modbus.TimeoutSeconds) * time.Second,
		MaxClients: uint(config.Modbus.MaxClients),
	}, handler)
	if err != nil {
		return nil, fmt.Errorf("failed to create modbus server. Err: %w", err)
	}

	return &Server{
		server:  server,
		handler: handler,
	}, nil
}
