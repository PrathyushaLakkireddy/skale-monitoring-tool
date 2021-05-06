package exporter

import (
	"log"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/monitor"
)

const (
	httpTimeout = 5 * time.Second
)

// metricsCollector respresents a set of skale metrics
type metricsCollector struct {
	config  *config.Config
	mutex   *sync.Mutex
	version *prometheus.Desc
	// totalValidatorsDesc       *prometheus.Desc
	// validatorActivatedStake   *prometheus.Desc
	// validatorLastVote         *prometheus.Desc
	// validatorRootSlot         *prometheus.Desc
	// validatorDelinquent       *prometheus.Desc
	// solanaVersion             *prometheus.Desc
	// accountBalance            *prometheus.Desc
	// slotLeader                *prometheus.Desc
	// blockTime                 *prometheus.Desc
	// currentSlot               *prometheus.Desc
	// commission                *prometheus.Desc
	// delinqentCommission       *prometheus.Desc
	// validatorVote             *prometheus.Desc
	// statusAlertCount          *prometheus.Desc
	// ipAddress                 *prometheus.Desc
	// txCount                   *prometheus.Desc
	// netVoteHeight             *prometheus.Desc
	// valVoteHeight             *prometheus.Desc
	// voteHeightDiff            *prometheus.Desc
	// valVotingStatus           *prometheus.Desc
	// voteCredits               *prometheus.Desc
	// networkConfirmationTime   *prometheus.Desc
	// validatorConfirmationTime *prometheus.Desc
	// confirmationTimeDiff      *prometheus.Desc
	// // confirmed block time of network
	// networkBlockTime *prometheus.Desc
	// // confirmed block time of validator
	// validatorBlockTime *prometheus.Desc
	// // block time difference of network and validator
	// blockTimeDiff *prometheus.Desc
}

func NewMetricsCollector(cfg *config.Config) *metricsCollector {
	return &metricsCollector{
		config: cfg,
		version: prometheus.NewDesc(
			"skale_version",
			"Current version of SKALE network client",
			[]string{"version"}, nil),
	}
}

// Desribe exports metrics to the channel
func (c *metricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.version
}

func (c *metricsCollector) Collect(ch chan<- prometheus.Metric) {
	log.Println("cmng here...")
	// get version
	// c.mutex.Lock()
	cVersion, err := monitor.GetClientVersion(c.config)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.version, err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.version, prometheus.GaugeValue, 1, cVersion.Result)
	}
	// c.mutex.Unlock()
}
