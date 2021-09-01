package monitor_test

import (
	"testing"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/monitor"
)

func TestGetPublicIP(t *testing.T) {
	cfg, err := config.ReadFromFile()
	if err != nil {
		t.Error("Error while reading config : ", err)
	}
	res, err := monitor.GetPublicIP(cfg)
	if err != nil {
		t.Error("Error while fetching Public IP")
	}
	if &res == nil {
		t.Error("Expected non empty result, but got empyt result: ", res)
	}
	if &res != nil {
		t.Log("Got Public IP: ", res)
	}
}
