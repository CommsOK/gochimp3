package gochimp3

const (
	ping_path = "/ping"
)

type PingResponse struct {
	HealthStatus string `json:"health_status"`
}

// Ping is a health check for the API that won't return any account-specific information.
func (api *API) Ping() (*PingResponse, error) {
	response := new(PingResponse)
	err := api.Request("GET", ping_path, nil, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
