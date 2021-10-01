package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetHardwareInfo which returns the hardware info of node
func GetHardwareInfo(cfg *config.Config) (types.Hardware, error) {
	log.Println("Getting Hardware Requirements...")

	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP + "/status/hardware",
		Method:   http.MethodGet,
	}

	var result types.Hardware
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
