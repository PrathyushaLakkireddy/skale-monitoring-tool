package exporter

import (
	"log"
	"strconv"
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
	config      *config.Config
	mutex       *sync.Mutex
	version     *prometheus.Desc
	sgxStatus   *prometheus.Desc
	publicIP    *prometheus.Desc
	coreStatus  *prometheus.Desc
	noOfschains *prometheus.Desc
	sChains     *prometheus.Desc
}

func NewMetricsCollector(cfg *config.Config) *metricsCollector {
	return &metricsCollector{
		config: cfg,
		version: prometheus.NewDesc(
			"skale_version",
			"Current version of SKALE network client",
			[]string{"version"}, nil),
		sgxStatus: prometheus.NewDesc(
			"skale_sgx_status",
			"Get sgx server info",
			[]string{"status_name", "wallet_version"}, nil),
		publicIP: prometheus.NewDesc(
			"skale_public_ip",
			"Publi IP of skale node, to which the packets were sending",
			[]string{"public_ip"}, nil),
		coreStatus: prometheus.NewDesc(
			"skale_core_status",
			"status about docker images",
			[]string{"image", "name", "status"}, nil),
		noOfschains: prometheus.NewDesc(
			"no_of_schains",
			"No of skale schains",
			[]string{"schains_count"}, nil),
		sChains: prometheus.NewDesc(
			"skale_schains",
			"info of schains",
			[]string{"name", "dkg_status"}, nil),
	}
}

// Desribe exports metrics to the channel
func (c *metricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.version
	ch <- c.sgxStatus
	ch <- c.publicIP
	ch <- c.coreStatus
	ch <- c.noOfschains
	ch <- c.sChains
}

func (c *metricsCollector) Collect(ch chan<- prometheus.Metric) {
	log.Println("cmng here...")
	// get version
	// c.mutex.Lock()
	cVersion, err := monitor.GetClientVersion(c.config) // TODO check
	if err != nil {
		log.Printf("Error while getting client version : %v", err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.version, prometheus.GaugeValue, 1, cVersion.Result)
	}
	// c.mutex.Unlock()

	// get sgx status info
	sgx, err := monitor.GetSGXStatus(c.config)
	if err != nil {
		log.Printf("Error while fetching sgx status : %v", err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.sgxStatus, prometheus.GaugeValue, float64(sgx.Data.Status), sgx.Data.StatusName, sgx.Data.SgxWalletVersion)
	}

	//get node's public ip
	pIP, err := monitor.GetPublicIP(c.config)
	if err != nil {
		log.Printf("Error while getting public ip : %v", err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.publicIP, prometheus.GaugeValue, 0, pIP.Data.IP)
	}

	// get core status to check docker images running status
	cStatus, err := monitor.GetCoreStatus(c.config)
	if err != nil {
		log.Printf("Error while getting core status : %v", err)
	} else {
		for _, v := range cStatus.Data {
			ch <- prometheus.MustNewConstMetric(c.coreStatus, prometheus.GaugeValue, -1, v.Image, v.Name, v.State.Status)
		}
	}

	// get schains info
	schains, err := monitor.GetSchainStatus(c.config)
	if err != nil {
		log.Printf("Error while getting schain status : %v", err)
	} else {
		n := len(schains.Data)
		ch <- prometheus.MustNewConstMetric(c.noOfschains, prometheus.GaugeValue, float64(n), strconv.Itoa(n))

		for _, s := range schains.Data {
			ch <- prometheus.MustNewConstMetric(c.sChains, prometheus.GaugeValue, -1, s.Name, strconv.FormatBool(s.Healthchecks.Dkg))
		}
	}
}
