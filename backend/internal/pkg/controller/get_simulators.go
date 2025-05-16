package controller

import (
	"encoding/json"
	"net/http"
	"sensor-simulator/internal/pkg/dto"
)

func (c *SimulatorController) getSimulators(w http.ResponseWriter, req *http.Request) {
	dtos := c.service.GetSimulatorDtos()

	rawData, err := json.Marshal(dto.GetSimulators{
		Simulators: dtos,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(rawData)
}
