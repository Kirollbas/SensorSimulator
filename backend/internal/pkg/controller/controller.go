package controller

import (
	"context"
	"net/http"
	"sensor-simulator/internal/pkg/dto"

	"github.com/gorilla/mux"
)

type SimulatorService interface {
	AddSensor(dto dto.Simulator) error
	DeleteSensor(name string) error
	GetSimulatorDtos() []dto.SimulatorWithStatus
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type SimulatorController struct {
	service SimulatorService
}

func NewSimulatorController(
	service SimulatorService,
) *SimulatorController {
	return &SimulatorController{
		service: service,
	}
}

func (c *SimulatorController) SetupRouter(router *mux.Router) {
	router.HandleFunc("/simulator", c.getSimulators).Methods(http.MethodGet)
	router.HandleFunc("/simulator/add", c.addSimulator).Methods(http.MethodPost)
	router.HandleFunc("/simulator/{name}", c.deleteSimulator).Methods(http.MethodDelete)
	router.HandleFunc("/start", c.start).Methods(http.MethodPost)
	router.HandleFunc("/stop", c.stop).Methods(http.MethodPost)
}
