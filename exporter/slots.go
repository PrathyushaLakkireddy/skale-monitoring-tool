package exporter

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/monitor"
)

const (
	slotPacerSchedule = 5 * time.Second
)

var (
	syncing = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_syncing",
		Help: "Node Syncing",
	})

	blockNumber = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_block_number",
		Help: "Skale block number",
	})
)

func init() {
	prometheus.MustRegister(syncing)
	prometheus.MustRegister(blockNumber)
}

// WatchSlots get data from different methods and store that data in prometheus. Metrics are
// Block Number and synching status
func (c *metricsCollector) WatchSlots(cfg *config.Config) {

	ticker := time.NewTicker(slotPacerSchedule)

	for {
		<-ticker.C

		// get skale node enpoint status
		status, err := monitor.GetEndpointStatus(cfg)
		if err != nil {
			log.Printf("Error while getting endpoint status : %v", err)
		} else {
			bn := float64(status.Data.BlockNumber)
			blockNumber.Set(bn)
		}
	}
}
