package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetSyncingStatus returns the syncing status of node
func GetSyncingStatus(cfg *config.Config) (float64, error) {
	log.Println("Getting Syncing status...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "eth_syncing", ID: 1},
	}

	var i float64 = 1

	var result types.Syncing
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return i, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return i, err
	}

	if result.Result {
		i = 0
	}

	log.Printf("Syncing status : %v and value : %f", result.Result, i)

	return i, nil
}
