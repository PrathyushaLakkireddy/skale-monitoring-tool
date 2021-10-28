package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetNodeInfo returns Skale Node Info metrics which are
// 1. Node Name
// 2. IP and Public IP
// 3. port and domain name
func GetNodeInfo(cfg *config.Config) (types.NodeInfo, error) {
	log.Println("Getting Node Info...")

	cmd := exec.Command("sudo", "skale", "node", "info", "-f", "json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while runnig skale validator cli command %v", err)
	}
	s := string(out)
	s = strings.Replace(s, "'", "\"", -1)
	var result types.NodeInfo
	err = json.Unmarshal([]byte(s), &result)
	if err != nil {
		log.Printf("Error:%v", err)
	}
	GetNode(cfg)
	return result, nil
}
func GetNode(cfg *config.Config) {
	log.Println("Getting Node Info...")

	cmd := exec.Command("sudo", "skale", "node", "info")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while runnig skale validator cli command %v", err)
	}
	// s := string(out)
	// s = strings.Replace(s, "'", "\"", -1)
	// var result types.NodeInfo
	// err = json.Unmarshal([]byte(s), &result)
	// if err != nil {
	// 	log.Printf("Error:%v", err)
	// }

	fmt.Println("??????????????????????????????", string(out))
	// return result, nil

}
