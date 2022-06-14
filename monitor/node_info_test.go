package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/monitor"
)

func TestGetNodeInfo(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config : ", err)
	}
	res, err := monitor.GetNodeInfo(cfg)
	if err != nil {
		t.Error("Error while fetching Node Info")
	}
	if &res == nil {
		t.Error("Expected non empty result, but got empyt result: ", res)
	}
	if &res != nil {
		t.Log("Got Node Info: ", res)
	}
}
