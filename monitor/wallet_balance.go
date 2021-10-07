package monitor

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

func GetWalletInfo(cfg *config.Config) (types.WalletInfo, error) {

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
