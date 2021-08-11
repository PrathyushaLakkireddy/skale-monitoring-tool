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

	balance = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_balance",
		Help: "Skale account balance",
	})

	peers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "skale_peers_count",
		Help: "Skale peers count",
	})
)

func init() {
	prometheus.MustRegister(syncing)
	prometheus.MustRegister(blockNumber)
	prometheus.MustRegister(balance)
	prometheus.MustRegister(peers)
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
