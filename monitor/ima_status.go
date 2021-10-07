package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetIMAstatus which returns the IMA status
func GetIMAstatus(cfg *config.Config) (types.IMAstatus, error) {
	log.Println("Getting the status of the IMA container...")

	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP + "/status/ima",
		Method:   http.MethodGet,
	}

	var result types.IMAstatus
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
