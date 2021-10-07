package monitor

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

func GetNodeInfo(cfg *config.Config) (types.NodeInfo, error) {

	cmd := exec.Command("skale", "node", "info", "-f", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while runnig skale validator cli command %v", err)
	}
	// out := string(s)
	// out = strings.Replace(out, "'", "\"", -1)

	var result types.NodeInfo
	err = json.Unmarshal(out, &result)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	return result, nil

}
