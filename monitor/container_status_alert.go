package monitor

import (
	"fmt"
	"log"
	"strings"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/alerter"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
)

func ContainerStatusAlert(cfg *config.Config) error {
	result, err := GetCoreStatus(cfg)
	if err != nil {
		log.Printf("Error while getting status core :%v", err)
	}
	for _, container := range result.Data {
		if strings.EqualFold(cfg.AlerterPreferences.ContainerHealthAlerts, "yes") {
			if container.State.Running == false {
				teleErr := alerter.SendTelegramAlert(fmt.Sprintf("%s container stopped running", container.Name), cfg)
				if teleErr != nil {
					log.Printf("Error while sending container health alert : %v", teleErr)
				}
			}
			if container.State.Paused == true {
				teleErr := alerter.SendTelegramAlert(fmt.Sprintf("%s container has paused", container.Name), cfg)
				if teleErr != nil {
					log.Printf("Error while sending container health alert : %v", teleErr)
				}
			}
			if container.State.Dead == true {
				teleErr := alerter.SendTelegramAlert(fmt.Sprintf("%s container state is dead", container.Name), cfg)
				if teleErr != nil {
					log.Printf("Error while sending container health alert : %v", teleErr)
				}
			}
		}
	}
	return nil
}
