package controller

import "net/http"

func (c *SimulatorController) stop(w http.ResponseWriter, req *http.Request) {
	err := c.service.Stop(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
