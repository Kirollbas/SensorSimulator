package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (c *SimulatorController) deleteSimulator(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	simulatorName := vars["name"]

	err := c.service.DeleteSensor(simulatorName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
