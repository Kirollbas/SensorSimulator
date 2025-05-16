package controller

import (
	"context"
	"net/http"
)

func (c *SimulatorController) start(w http.ResponseWriter, req *http.Request) {
	err := c.service.Start(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
