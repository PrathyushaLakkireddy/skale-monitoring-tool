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

	balance = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_balance",
		Help: "Skale account balance",
	})

	peers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_peers_count",
		Help: "Skale peers count",
	})

	alertCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_alert_count",
		Help: "Skale Alert count",
	})
)

func init() {
	prometheus.MustRegister(syncing)
	prometheus.MustRegister(blockNumber)
	prometheus.MustRegister(balance)
	prometheus.MustRegister(peers)
	prometheus.MustRegister(alertCount)
}

func (c *metricsCollector) WatchSlots(cfg *config.Config) {

	ticker := time.NewTicker(slotPacerSchedule)

	for {
		<-ticker.C

		log.Printf("here....")

		status, err := monitor.GetEndpointStatus(cfg)
		if err != nil {
			log.Printf("Error while getting endpoint status : %v", err)
		} else {
			bn := float64(status.Data.BlockNumber)
			blockNumber.Set(bn)

			i := 1
			if status.Data.Syncing {
				i = 0
				if strings.EqualFold(cfg.AlerterPreferences.BlockSyncAlerts, "yes") {
					telegramErr := alerter.SendTelegramAlert(fmt.Sprintf("Current block is in Syncing Process"), cfg)
					if telegramErr != nil {
						log.Printf("Error while sending block syncing status alert to telegram : %v", telegramErr)
					}
					emailErr := alerter.SendEmailAlert(fmt.Sprintf("Current block is in Syncing Process"), cfg)
					if emailErr != nil {
						log.Printf("Error while block syncing status alert to Email : %v", emailErr)
					}
				}
			}
			syncing.Set(float64(i))
		}

		bal, err := monitor.GetBalance(cfg)
		if err != nil {
			log.Printf("Error while getting account bal : %v", err)
		}

		balance.Set(bal)

		p, err := monitor.GetPeersCount(cfg)
		if err != nil {
			log.Printf("Error while getting peers count : %v", err)
		}

		peers.Set(p)
	}
}
