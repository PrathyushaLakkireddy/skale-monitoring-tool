package monitor

import (
	"fmt"
	"log"
	"strings"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/alerter"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
)

func NodeStatusAlert(cfg *config.Config) error {

	info, err := GetNodeInfo(cfg)
	if err != nil {
		log.Printf("Error while getting node info: %v", err)
	}
	status := info.Status
	st := map[int]string{
		0: "Everything is OK",
		1: "General error exit code",
		3: "Bad API response",
		4: "Script execution error",
		5: "Transaction error",
		6: "Revert error",
		7: "Bad user error",
		8: "Node state error",
	}

	if status != 0 {
		// send telegram and email alerts
		if strings.EqualFold(cfg.AlerterPreferences.NodeHealthAlert, "yes") {
			teleErr := alerter.SendTelegramAlert(fmt.Sprintf("Node Health Alert: %s", st[status]), cfg)
			if teleErr != nil {
				log.Printf("Error while sending node health alert: %v", teleErr)
			}
			emailErr := alerter.SendEmailAlert(fmt.Sprintf("Node Health Alert: %s", st[status]), cfg)
			if emailErr != nil {
				log.Printf("Error while sending node health alert: %v", teleErr)
			}
		}
	}
	return nil
}
