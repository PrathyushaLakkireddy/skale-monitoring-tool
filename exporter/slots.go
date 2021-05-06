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

		num, err := monitor.GetBlockNumber(cfg)
		if err != nil {
			log.Printf("Error while getting block number : %v", err)
		}
		blockNumber.Set(num)

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
