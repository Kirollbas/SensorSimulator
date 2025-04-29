package app

import (
	"fmt"
	"log"
	"net"
	pb "sensor-simulator/gen/sensor_simulator/proto/simulator"
	"sensor-simulator/internal/pkg/endpoint/modbus"
	"sensor-simulator/internal/pkg/endpoint/opcua"
	"sensor-simulator/internal/pkg/service"

	"google.golang.org/grpc"
)

func App() {
	modbusServer, err := modbus.NewServer()
	if err != nil {
		log.Fatalf("unable to create modbus server: %v", err)
	}

	opcuaServer, err := opcua.NewServer()
	if err != nil {
		log.Fatalf("unable to create OPC UA server: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка запуска: %v", err)
	}

	grpcServer := grpc.NewServer()
	simulatorServer := service.NewSimulatorService(
		modbusServer,
		opcuaServer,
	)
	pb.RegisterSensorSimulatorServiceServer(grpcServer, simulatorServer)

	fmt.Println("gRPC сервер запущен на порту 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
