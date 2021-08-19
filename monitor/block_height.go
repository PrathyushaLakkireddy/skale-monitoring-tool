package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/utils"
)

// GetBlockNumber returns the current number of skale network block
func GetBlockNumber(cfg *config.Config) (float64, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "eth_blockNumber", ID: 1},
	}

	var blockNumber float64

	var result types.EthResult
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return blockNumber, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return blockNumber, err
	}

	num, err := utils.HexToIntConversion(result.Result)
	if err != nil {
		log.Printf("Error while converting block hex number to int: %v", err)
	}

	log.Printf("Block number : %d and %f", num, float64(num))

	return float64(num), nil
}
