package exporter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/alerter"
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

	alertCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_alert_count",
		Help: "Skale Alert count",
	})
)

func init() {
	prometheus.MustRegister(syncing)
	prometheus.MustRegister(blockNumber)
	prometheus.MustRegister(alertCount)
}

// WatchSlots get data from different methods and store that data in prometheus. Metrics are
// 1. Block Number and synching status
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

			s := status.Data.Syncing
			if s == true {
				if strings.EqualFold(cfg.AlerterPreferences.BlockSyncAlerts, "yes") {
					// send alert if the block synching is in process
					telegramErr := alerter.SendTelegramAlert(fmt.Sprintf("Current block is in Syncing Process"), cfg)
					if telegramErr != nil {
						log.Printf("Error while sending block syncing status alert to telegram : %v", telegramErr)
					}
					emailErr := alerter.SendEmailAlert(fmt.Sprintf("Current block is in Syncing Process"), cfg)
					if emailErr != nil {
						log.Printf("Error while sending block syncing status alert to Email : %v", emailErr)
					}
				}
				syncing.Set(float64(1))
			} else {
				syncing.Set(float64(0))
			}
		}
	}
}
