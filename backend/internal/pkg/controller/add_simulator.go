package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"sensor-simulator/internal/pkg/dto"
)

func (c *SimulatorController) addSimulator(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	if !json.Valid(body) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var simulatorDTO dto.Simulator
	err = json.Unmarshal(body, &simulatorDTO)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.service.AddSensor(simulatorDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
