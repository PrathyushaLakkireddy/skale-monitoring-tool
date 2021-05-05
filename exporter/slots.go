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
)

func init() {
	prometheus.MustRegister(syncing)
}

func (c *metricsCollector) WatchSlots(cfg *config.Config) {
	var (
	// Current mapping of relative slot numbers to leader public keys.
	// epochSlots map[int64]string
	// // Current epoch number corresponding to epochSlots.
	// epochNumber int64
	// // Last slot number we generated ticks for.
	// watermark int64
	)

	ticker := time.NewTicker(slotPacerSchedule)

	for {
		<-ticker.C

		log.Printf("here....")
		sync, err := monitor.GetSyncingStatus(cfg)
		if err != nil {
			log.Printf("Error while getting syncing status : %v", err)
		}
		syncing.Set(sync)

		// ch := chan<- prometheus.Metric
		// c.Collect(make(chan<- prometheus.Metric))
	}
}
