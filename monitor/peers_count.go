package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/alerter"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/utils"
)

// GetPeersCount returns peers count
func GetPeersCount(cfg *config.Config) (float64, error) {
	log.Println("Getting Peers Count...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP,
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

	err = SendPeersCountAlert(int64(num), cfg)
	log.Printf("Peers Count : %f", float64(num))

	return float64(num), nil
}

func SendPeersCountAlert(peers int64, cfg *config.Config) error {

	if strings.EqualFold(cfg.AlerterPreferences.NumPeersAlerts, "yes") {
		if peers < cfg.AlertingThresholds.NumPeersThreshold {
			telegramErr := alerter.SendTelegramAlert(fmt.Sprintf("Number of connected peers dropped below configured threshold,current connected peers are %v", peers), cfg)
			if telegramErr != nil {
				log.Printf("Error while sending number of connected peers alert to telegram : %v", telegramErr)
			}
			emailErr := alerter.SendEmailAlert(fmt.Sprintf("Number of connected peers dropped below configured threshold,current connected peers are %v", peers), cfg)
			if emailErr != nil {
				log.Printf("Error while sending number of connected peers alert to Email : %v", emailErr)
			}
		}
	}

	return nil
}
