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
)

// GetCoreStatus returns the core status of skale network
func GetCoreStatus(cfg *config.Config) (types.StatusCore, error) {
	log.Println("Getting Core Status...")
	ops := types.HTTPOptions{
		Endpoint: cfg.Endpoints.SkaleNodeIP + "/status/core",
		Method:   http.MethodGet,
	}

	var result types.StatusCore
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

	for _, container := range result.Data {
		if strings.EqualFold(cfg.AlerterPreferences.ContainerHealthAlerts, "yes") {
			if container.State.Running == false {
				teleErr := alerter.SendTelegramAlert(fmt.Sprintf("Your %s container stopped running", container.Name), cfg)
				if teleErr != nil {
					log.Printf("Error while sending container health alert : %v", teleErr)
				}
			}
			if container.State.Paused == true {
				teleErr := alerter.SendTelegramAlert(fmt.Sprintf("Your %s container has paused", container.Name), cfg)
				if teleErr != nil {
					log.Printf("Error while sending container health alert : %v", teleErr)
				}
			}
			if container.State.Dead == true {
				teleErr := alerter.SendTelegramAlert(fmt.Sprintf("Your %s container state is dead", container.Name), cfg)
				if teleErr != nil {
					log.Printf("Error while sending container health alert : %v", teleErr)
				}
			}
		}

	}
	return result, nil

}
