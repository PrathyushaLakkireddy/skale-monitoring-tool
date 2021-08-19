package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetClientVersion returns the current solana versions running on the node
func GetClientVersion(cfg *config.Config) (types.EthResult, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "web3_clientVersion", ID: 1},
	}

	var result types.EthResult
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

	log.Println("version..", result.Result)

	return result, nil
}
