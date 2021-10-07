package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetSGXStatus returns SGX server info - connection status and SGX wallet version
func GetSGXStatus(cfg *config.Config) (types.SGXStatus, error) {
	log.Println("Getting SGX Status...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP + "/status/sgx",
		Method:   http.MethodGet,
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
