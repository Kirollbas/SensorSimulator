package app

import (
	"fmt"
	"log"
	"net/http"
	"sensor-simulator/internal/configs"
	"sensor-simulator/internal/pkg/controller"
	"sensor-simulator/internal/pkg/endpoint/modbus"
	"sensor-simulator/internal/pkg/endpoint/opcua"
	"sensor-simulator/internal/pkg/middleware"
	"sensor-simulator/internal/pkg/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Printf("No .env file found %s\n", err)
	}
}

func App() {
	modbusServer, err := modbus.NewServer()
	if err != nil {
		log.Fatalf("unable to create modbus server: %v", err)
	}

	opcuaServer, err := opcua.NewServer()
	if err != nil {
		log.Fatalf("unable to create OPC UA server: %v", err)
	}

	simulatorServer := service.NewSimulatorService(
		modbusServer,
		opcuaServer,
	)

	simulatorController := controller.NewSimulatorController(
		simulatorServer,
	)

	router := mux.NewRouter()

	router.Use(middleware.EnableCORS)

	simulatorRouter := router.PathPrefix("/api").Subrouter()
	simulatorController.SetupRouter(simulatorRouter)

	config := configs.GetConfig()
	addressString := fmt.Sprintf("0.0.0.0:%d", config.Simulator.Port)

	err = http.ListenAndServe(addressString, router)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to start server: %w", err))
		return
	}
}
