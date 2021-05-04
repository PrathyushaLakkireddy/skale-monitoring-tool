package exporter

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
)

const (
	httpTimeout = 5 * time.Second
)

// solanaCollector respresents a set of solana metrics
type metricsCollector struct {
	config *config.Config
	// version
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
	}
}

// Desribe exports metrics to the channel
func (c *metricsCollector) Describe(ch chan<- *prometheus.Desc) {}

func (c *metricsCollector) Collect(ch chan<- prometheus.Metric) {
}
