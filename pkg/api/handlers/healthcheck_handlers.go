package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/leoguilen/simple-go-api/pkg/api/models/responses"
	"github.com/leoguilen/simple-go-api/pkg/infra/repositories"
)

var (
	hcRepository = repositories.NewHealthcheckRepository()
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := hcRepository.Check(r.Context()); err != nil {
		res := models.NewErrorResponseFromError(err)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(res)
	}

	res := models.NewHealthcheckResponse()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
