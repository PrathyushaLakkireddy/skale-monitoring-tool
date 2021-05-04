package exporter

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/monitor"
)

const (
	slotPacerSchedule = 500 * time.Millisecond
)

var (
	version = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_verion",
		Help: "Skale version",
	})
)

func init() {
	prometheus.MustRegister(version)
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

		_, _ = monitor.GetVersion(cfg)
	}
}
