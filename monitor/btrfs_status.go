package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetBTRFSStatus returns BTRFS kernal information
func GetBTRFSstatus(cfg *config.Config) (types.BTRFSstatus, error) {
	log.Println("Getting BTRFS status ...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP + "/status/btrfs",
		Method:   http.MethodGet,
	}

	var result types.BTRFSstatus
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
