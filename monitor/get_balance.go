package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/utils"
)

// GetBalance returns the balance of
func GetBalance(cfg *config.Config) (float64, error) {
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP,
		Method:   http.MethodPost,
		Body:     types.Payload{Jsonrpc: "2.0", Method: "eth_getBalance", Params: []interface{}{"0x407d73d8a49eeb85d32cf465507dd71d507100c1", "latest"}, ID: 1},
	}

	var bal float64

	var result types.EthResult
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return bal, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return bal, err
	}

	b, err := utils.HexToIntConversion(result.Result)
	if err != nil {
		log.Printf("Error while converting block hex number to int: %v", err)
	}

	log.Printf("Balance : %d and %f", b, float64(b))

	return float64(b), nil
}
