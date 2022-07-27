package models

type HealthcheckResponse struct {
	Status string `json:"status"`
	Detail string `json:"detail"`
}

func NewHealthcheckResponse() *HealthcheckResponse {
	return &HealthcheckResponse{
		Status: "Healthy",
		Detail: "Service is available.",
	}
}

func NewHealthcheckResponseFromError(err error) *HealthcheckResponse {
	return &HealthcheckResponse{
		Status: "Unhealthy",
		Detail: err.Error(),
	}
}
