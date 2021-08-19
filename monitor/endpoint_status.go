package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetEndpointStatus returns the status of the endpoint
func GetEndpointStatus(cfg *config.Config) (types.EndpointStatus, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP + "/status/endpoint",
		Method:   http.MethodPost,
	}

	var result types.EndpointStatus
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	return result, nil
}
