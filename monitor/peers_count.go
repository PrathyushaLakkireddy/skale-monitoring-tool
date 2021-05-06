package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/utils"
)

// GetPeersCount returns peers count
func GetPeersCount(cfg *config.Config) (float64, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.RPCEndpoint,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "net_peerCount", ID: 1},
	}

	var v float64

	var result types.EthResult
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return v, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return v, err
	}

	num, err := utils.HexToIntConversion(result.Result)
	if err != nil {
		log.Printf("Error while converting block hex number to int: %v", err)
	}

	log.Printf("Peers Count : %f", float64(num))

	return float64(num), nil
}
