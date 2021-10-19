package exporter

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/alerter"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/monitor"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/querier"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/utils"
)

const (
	httpTimeout = 5 * time.Second
)

// metricsCollector respresents a set of skale metrics
type metricsCollector struct {
	config        *config.Config
	mutex         *sync.Mutex
	version       *prometheus.Desc
	sgxStatus     *prometheus.Desc
	publicIP      *prometheus.Desc
	coreStatus    *prometheus.Desc
	noOfschains   *prometheus.Desc
	sChains       *prometheus.Desc
	hardware      *prometheus.Desc
	btrfs         *prometheus.Desc
	nodeName      *prometheus.Desc
	ip            *prometheus.Desc
	port          *prometheus.Desc
	status        *prometheus.Desc
	nodeID        *prometheus.Desc
	domainName    *prometheus.Desc
	walletAddress *prometheus.Desc
	ethBalance    *prometheus.Desc
	skaleBalance  *prometheus.Desc
	conAlertCount *prometheus.Desc
}

// NewMetricsCollector exports metricsCollector metrics to prometheus
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
		hardware: prometheus.NewDesc(
			"hardware_info",
			"hardware information of skale node",
			[]string{"cpu_total_cores", "cpu_physical_cores", "memory", "swap",
				"system_release", "uname_version", "attached_storage_size"}, nil),
		btrfs: prometheus.NewDesc(
			"btrfs_status",
			"BTRFS kernal module status",
			[]string{"btrfs_status"}, nil),
		nodeName: prometheus.NewDesc(
			"skale_node_name",
			"Skale node name",
			[]string{"node_name"}, nil),
		ip: prometheus.NewDesc(
			"skale_ip",
			"Skale node IP",
			[]string{"skale_ip"}, nil),
		port: prometheus.NewDesc(
			"skale_node_port",
			"Skale Node port",
			[]string{"skale_port"}, nil),
		status: prometheus.NewDesc(
			"skale_node_status",
			"Skale node status",
			[]string{"skale_node_status"}, nil),
		nodeID: prometheus.NewDesc(
			"skale_node_id",
			"Skale node id",
			[]string{"skale_node_id"}, nil),
		domainName: prometheus.NewDesc(
			"skale_domain_name",
			"skale node domain name",
			[]string{"skale_domain_name"}, nil),
		walletAddress: prometheus.NewDesc(
			"skale_wallet_address",
			"Skale wallet address",
			[]string{"skale_wallet_address"}, nil),
		ethBalance: prometheus.NewDesc(
			"skale_eth_balance",
			"Skale ETH balance",
			[]string{"skale_eth_balance"}, nil),
		skaleBalance: prometheus.NewDesc(
			"skale_balance",
			"Skale balance",
			[]string{"skale_balance"}, nil),
		conAlertCount: prometheus.NewDesc(
			"skale_con_alertCount",
			"skale_con_alertCount",
			[]string{"skale_con_alertCount"}, nil),
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
	ch <- c.hardware
	ch <- c.btrfs
	ch <- c.nodeName
	ch <- c.ip
	ch <- c.port
	ch <- c.status
	ch <- c.nodeID
	ch <- c.domainName
	ch <- c.walletAddress
	ch <- c.ethBalance
	ch <- c.skaleBalance
}

// Collect get data from methods and exports metrics to prometheus. Metrics are
// 1. SGX status
// 2. Public IP
// 3. Core status or status of containers
// 4. Schain Status
// 5. Hardware Info
// 6. BTRFS status
// 7. Skalr Node Info which includes node name, status, port, ip, public Ip etc..
// 8. Skale Wallet Info which includes wallet addres, ETH balance, SKL balance
func (c *metricsCollector) Collect(ch chan<- prometheus.Metric) {
	log.Println("Collecting exporter metrics...")
	// get version
	// c.mutex.Lock()
	// c.mutex.Unlock()

	// get sgx status info
	sgx, err := monitor.GetSGXStatus(c.config)
	if err != nil {
		log.Printf("Error while fetching sgx status : %v", err)
	} else {
		ch <- prometheus.MustNewConstMetric(c.sgxStatus, prometheus.GaugeValue, float64(sgx.Data.Status), sgx.Data.StatusName, sgx.Data.SgxWalletVersion)
	}

	// //get node's public ip
	// pIP, err := monitor.GetPublicIP(c.config)
	// if err != nil {
	// 	log.Printf("Error while getting public ip : %v", err)
	// } else {
	// 	ch <- prometheus.MustNewConstMetric(c.publicIP, prometheus.GaugeValue, 0, pIP.Data.PublicIP)
	// }

	// get core status to check docker images running status
	cStatus, err := monitor.GetCoreStatus(c.config)
	cl := len(cStatus.Data)
	if err != nil {
		log.Printf("Error while getting core status : %v", err)
	} else {
		var run int64
		var pas int64
		var dead int64
		for _, v := range cStatus.Data {
			if v.State.Running == true {
				run = run + 1
			} else if v.State.Paused == true {
				pas = pas + 1
			} else if v.State.Dead == true {
				dead = dead + 1
			}
			ch <- prometheus.MustNewConstMetric(c.coreStatus, prometheus.GaugeValue, -1, v.Image, v.Name, v.State.Status)
		}
		c.AlertContainerStaus(run, pas, dead, cl, ch)
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

	// get hardware info
	h, err := monitor.GetHardwareInfo(c.config)
	if err != nil {
		log.Printf("Error while getting hardware information : %v", err)
	} else {
		at := utils.LenReadable(int(h.Data.AttachedStorageSize), 2)
		m := utils.LenReadable(int(h.Data.Memory), 2)
		s := utils.LenReadable(int(h.Data.Swap), 2)
		ch <- prometheus.MustNewConstMetric(c.hardware, prometheus.GaugeValue, -1, fmt.Sprintf("%d", h.Data.CPUTotalCores), fmt.Sprintf("%d", h.Data.CPUPhysicalCores), m,
			s, h.Data.SystemRelease, h.Data.UnameVersion, at)
	}

	// get btrfs kernal module info
	b, err := monitor.GetBTRFSstatus(c.config)
	if err != nil {
		log.Printf("Error while getting btrfs status : %v", err)
	} else {
		if b.Data.KernelModule == true {
			ks := "Enabled"
			ch <- prometheus.MustNewConstMetric(c.btrfs, prometheus.GaugeValue, -1, ks)
		} else {
			ks := "Disabled"
			ch <- prometheus.MustNewConstMetric(c.btrfs, prometheus.GaugeValue, -1, ks)
		}

	}

	// get skale node Info
	n, err := monitor.GetNodeInfo(c.config)
	if err != nil {
		log.Printf("Error while getting node info : %v", err)
	} else {
		nodeName := n.Name
		ch <- prometheus.MustNewConstMetric(c.nodeName, prometheus.GaugeValue, -1, nodeName)

		pIp := n.Publicip
		ch <- prometheus.MustNewConstMetric(c.publicIP, prometheus.GaugeValue, 1, pIp)

		ip := n.IP
		ch <- prometheus.MustNewConstMetric(c.ip, prometheus.GaugeValue, 1, ip)

		port := n.Port
		ch <- prometheus.MustNewConstMetric(c.port, prometheus.GaugeValue, float64(port), "skale port")

		s := n.Status
		ch <- prometheus.MustNewConstMetric(c.status, prometheus.GaugeValue, float64(s), "Skale node status")

		id := n.ID
		ch <- prometheus.MustNewConstMetric(c.nodeID, prometheus.GaugeValue, float64(id), "skale node id")

		dn := n.DomainName
		if dn != "" {
			ch <- prometheus.MustNewConstMetric(c.domainName, prometheus.GaugeValue, 1, dn)
		}

	}

	// get skale wallet info
	wInfo, err := monitor.GetWalletInfo(c.config)
	if err != nil {
		log.Printf("Error while getting skale wallet info : %v", err)
	} else {
		wa := wInfo.Address
		ch <- prometheus.MustNewConstMetric(c.walletAddress, prometheus.GaugeValue, 1, wa)

		eb := wInfo.EthBalanceWei
		eth := float64(eb) / math.Pow(10, 18)
		ch <- prometheus.MustNewConstMetric(c.ethBalance, prometheus.GaugeValue, eth, "skale ETH balance")

		sb := wInfo.SkaleBalanceWei
		skl := float64(sb) / math.Pow(10, 18)
		ch <- prometheus.MustNewConstMetric(c.skaleBalance, prometheus.GaugeValue, skl, "skale balance")

	}

}

func (c *metricsCollector) AlertContainerStaus(run int64, pas int64, dead int64, cl int, ch chan<- prometheus.Metric) {
	now := time.Now().UTC()
	currentTime := now.Format(time.Kitchen)

	var alertsArray []string

	for _, value := range c.config.RegularStatusAlerts.AlertTimings {
		t, _ := time.Parse(time.Kitchen, value)
		alertTime := t.Format(time.Kitchen)

		alertsArray = append(alertsArray, alertTime)
	}
	log.Printf("Current time : %v and alerts array : %v", currentTime, alertsArray)

	var count float64 = 0

	for _, conAlertTime := range alertsArray {
		if currentTime == conAlertTime {
			alreadySentAlert, _ := querier.ConAlertStatusCountFromPrometheus(c.config)
			if alreadySentAlert == "false" {
				if run == int64(cl) {
					telegramErr := alerter.SendTelegramAlert(fmt.Sprintf("Container Status Alert: All containers are running, Total Containers are %v", cl), c.config)
					if telegramErr != nil {
						log.Printf("Error while sending regular status alert: %v", telegramErr)
					}
					emialErr := alerter.SendEmailAlert(fmt.Sprintf("Container Status Alert: All containers are running, Total Containers are %v", cl), c.config)
					if emialErr != nil {
						log.Printf("Error while sending regular status alert: %v", emialErr)
					}
				} else {
					teleErr := alerter.SendTelegramAlert(fmt.Sprintf("Container Status Alert: %v Containers are RUNNING\n %v Containers are PAUSED \n,%v Containers are STOPPED", run, pas, dead), c.config)
					if teleErr != nil {
						log.Printf("Error while sending regular status alert: %v", teleErr)
					}
					emailErr := alerter.SendEmailAlert(fmt.Sprintf("Container Status Alert: %v Containers are RUNNING\n %v Containers are PAUSED \n,%v Containers are STOPPED", run, pas, dead), c.config)
					if emailErr != nil {
						log.Printf("Error while sending regular status alert: %v", teleErr)
					}

				}
				ch <- prometheus.MustNewConstMetric(c.conAlertCount, prometheus.GaugeValue, count, "true")
				count = count + 1
			} else {
				ch <- prometheus.MustNewConstMetric(c.conAlertCount, prometheus.GaugeValue, count, "false")
				return
			}
		}
	}

}
