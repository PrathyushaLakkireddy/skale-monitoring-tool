package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetCoreSttaus returns the core status of skale network
func GetCoreStatus(cfg *config.Config) (types.StatusCore, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint + "/status/core",
		Method:   http.MethodPost,
	}

	var result types.StatusCore
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
