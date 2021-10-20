package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/alerter"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetWalletInfo returns wallet address ETH and skale balances
func GetWalletInfo(cfg *config.Config) (types.WalletInfo, error) {
	log.Println("Getting Wallet Info...")
	cmd := exec.Command("skale", "wallet", "info", "-f", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while runnig skale validator cli command %v", err)
	}

	var result types.WalletInfo
	err = json.Unmarshal(out, &result)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	return result, nil

}

// SendBalanceChangeAlert sends ETH and Skale balance change and delegation alerts
func SendBalanceChangeAlert(ethBal string, sklBal string, cfg *config.Config) error {

	eb, _ := strconv.ParseFloat(ethBal, 64)
	sb, _ := strconv.ParseFloat(sklBal, 64)
	ce := ethBal + "ETH"
	cs := sklBal + "SKL"

	// pe, err := querier.GetETHbalanceFromDB(cfg)
	// if err != nil {
	// 	log.Printf("Error while getting ETH balance from DB : %v", err)
	// }
	// ps, err := querier.GetSKLbalanceFromDB(cfg)
	// if err != nil {
	// 	log.Printf("Error while getting skale balance from DB : %v", err)
	// }
	if strings.EqualFold(cfg.AlerterPreferences.EthbalanceChangeAlerts, "yes") {
		if eb < cfg.AlertingThresholds.EthbalanceThreshold {
			err := alerter.SendTelegramAlert(fmt.Sprintf("Account Balance Alert: Your ETH balance has dropped below configuration threshold, current balance is : %s", ce), cfg)
			if err != nil {
				log.Printf("Error while sending account balance change alert to telegram : %v", err)
				return err
			}

			err = alerter.SendEmailAlert(fmt.Sprintf("Account Balance Alert: Your ETH balance has dropped below configuration threshold, current balance is : %s", ce), cfg)
			if err != nil {
				log.Printf("Error while sending account balance change alert to Email : %v", err)
				return err
			}

		}

	}

	if strings.EqualFold(cfg.AlerterPreferences.SklbalanceChangeAlerts, "yes") {
		if sb < cfg.AlertingThresholds.SklbalanceThreshold {
			err := alerter.SendTelegramAlert(fmt.Sprintf("Account Balance Alert: Your SKALE account balance has dropped below configuration threshold, current balance is : %s", cs), cfg)
			if err != nil {
				log.Printf("Error while sending account balance change alert to telegram : %v", err)
				return err
			}

			err = alerter.SendEmailAlert(fmt.Sprintf("Account Balance Alert: Your SKALE account balance has dropped below configuration threshold, current balance is : %s", cs), cfg)
			if err != nil {
				log.Printf("Error while sending account balance change alert to Email : %v", err)
				return err
			}

		}
	}

	// if pe != "" {
	// 	peth, err := strconv.ParseFloat(pe, 64)
	// 	if err != nil {
	// 		log.Printf("Error while converting Eth balance to float : %v", err)
	// 	}
	// 	if strings.EqualFold(cfg.AlerterPreferences.ETHDelegationAlerts, "yes") {
	// 		if eb > peth {
	// 			err = alerter.SendTelegramAlert(fmt.Sprintf("ETH Delegation Alert: Your ETH balance has changed from %s to %s", pe+"ETH", ce), cfg)
	// 			if err != nil {
	// 				log.Printf("Error while sending ETH balance delegation alert to telegram : %v", err)
	// 				return err
	// 			}
	// 			err = alerter.SendEmailAlert(fmt.Sprintf("ETH Delegation Alert: Your ETH balance has changed from %s to %s", pe+"ETH", ce), cfg)
	// 			if err != nil {
	// 				log.Printf("Error while sending ETH balance delegation alert to email : %v", err)
	// 				return err
	// 			}
	// 		}
	// 		if eb < peth {
	// 			err = alerter.SendTelegramAlert(fmt.Sprintf("ETH UnDelegation Alert: Your ETH balance has changed from %s to %s", pe+"ETH", ce), cfg)
	// 			if err != nil {
	// 				log.Printf("Error while sending ETH balance undelegation alert to telegram : %v", err)
	// 				return err
	// 			}
	// 			err = alerter.SendEmailAlert(fmt.Sprintf("ETH UnDelegation Alert: Your ETH balance has changed from %s to %s", pe+"ETH", ce), cfg)
	// 			if err != nil {
	// 				log.Printf("Error while sending ETH balance undelegation alert to email : %v", err)
	// 				return err
	// 			}
	// 		}

	// 	}

	// }
	// if ps != "" {
	// 	pskl, err := strconv.ParseFloat(ps, 64)
	// 	if err != nil {
	// 		log.Printf("Error while converting skl balance to float : %v", err)
	// 	}
	// 	if strings.EqualFold(cfg.AlerterPreferences.SKLDelegationAlerts, "yes") {
	// 		if sb > pskl {
	// 			err = alerter.SendTelegramAlert(fmt.Sprintf("Skale Delegation Alert: Your skale balance has changed from %s to %s", ps+"SKL", cs), cfg)
	// 			if err != nil {
	// 				log.Printf("Error while sending skale balance delegation alert to telegram : %v", err)
	// 				return err
	// 			}
	// 			err = alerter.SendEmailAlert(fmt.Sprintf("skale Delegation Alert: Your skale balance has changed from %s to %s", ps+"SKL", cs), cfg)
	// 			if err != nil {
	// 				log.Printf("Error while sending skale balance delegation alert to email : %v", err)
	// 				return err
	// 			}
	// 		}
	// 		if eb < pskl {
	// 			err = alerter.SendTelegramAlert(fmt.Sprintf("Skale UnDelegation Alert: Your skale balance has changed from %s to %s", ps+"SKL", cs), cfg)
	// 			if err != nil {
	// 				log.Printf("Error while sending skale balance undelegation alert to telegram : %v", err)
	// 				return err
	// 			}
	// 			err = alerter.SendEmailAlert(fmt.Sprintf("Skale UnDelegation Alert: Your skale balance has changed from %s to %s", ps+"SKL", cs), cfg)
	// 			if err != nil {
	// 				log.Printf("Error while sending ETH balance undelegation alert to email : %v", err)
	// 				return err
	// 			}
	// 		}

	// 	}

	return nil

}
