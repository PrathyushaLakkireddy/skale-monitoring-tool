package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetCoreSttaus returns the core status of skale network
func GetSGXStatus(cfg *config.Config) (types.SGXStatus, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint + "/status/sgx",
		Method:   http.MethodPost,
	}

	var result types.SGXStatus
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
