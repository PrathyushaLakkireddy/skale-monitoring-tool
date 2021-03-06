package types

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
)

type (
	// QueryParams map of strings
	QueryParams map[string]string

	// HTTPOptions is a structure that holds all http options parameters
	HTTPOptions struct {
		Endpoint    string
		QueryParams QueryParams
		Body        Payload
		Method      string
	}

	// Payload is a structure which holds all the curl payload parameters
	Payload struct {
		Jsonrpc string        `json:"jsonrpc"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
		ID      int           `json:"id"`
	}

	// Commitement struct holds the state of Commitment
	Commitment struct {
		Commitemnt string `json:"commitment"`
	}
	// Encode struct to encode string
	Encode struct {
		Encoding string `json:"encoding"`
	}

	// Params struct
	Params struct {
		To     string `json:"to"`
		Data   string `json:"data"`
		Encode Encode `json:"encode"`
	}

	// Target is a structure which holds all the parameters of a target
	//this could be used to write endpoints for each functionality
	Target struct {
		ExecutionType string
		HTTPOptions   HTTPOptions
		Name          string
		Func          func(m HTTPOptions, cfg *config.Config, c client.Client)
		ScraperRate   string
	}

	// Targets list of all the targets
	Targets struct {
		List []Target
	}

	// PingResp is a structure which holds the options of a response
	PingResp struct {
		StatusCode int
		Body       []byte
	}

	// Syncing status of the node
	Syncing struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  bool   `json:"result"`
	}

	// // StatusCore is a struct which holds inoformation of image and status of it
	// StatusCore struct {
	// 	Data []struct {
	// 		Image string `json:"image"`
	// 		Name  string `json:"name"`
	// 		State struct {
	// 			Status     string `json:"Status"`
	// 			Running    bool   `json:"Running"`
	// 			Paused     bool   `json:"Paused"`
	// 			Restarting bool   `json:"Restarting"`
	// 			OOMKilled  bool   `json:"OOMKilled"`
	// 			Dead       bool   `json:"Dead"`
	// 			Pid        int    `json:"Pid"`
	// 			ExitCode   int    `json:"ExitCode"`
	// 			Error      string `json:"Error"`
	// 			StartedAt  string `json:"StartedAt"`
	// 			FinishedAt string `json:"FinishedAt"`
	// 		} `json:"state,omitempty"`
	// 	} `json:"data"`
	// 	Error interface{} `json:"error"`
	// }

	// SGXStatus which holds server info, connection status and SGX wallet version
	SGXStatus struct {
		Data struct {
			Status           int    `json:"status"`
			StatusName       string `json:"status_name"`
			SgxWalletVersion string `json:"sgx_wallet_version"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	// SchainsStatus which holds info of schains status
	SchainsStatus struct {
		Data []struct {
			Name         string `json:"name"`
			Healthchecks struct {
				DataDir       bool `json:"data_dir"`
				Dkg           bool `json:"dkg"`
				Config        bool `json:"config"`
				Volume        bool `json:"volume"`
				FirewallRules bool `json:"firewall_rules"`
				Container     bool `json:"container"`
				ExitCodeOk    bool `json:"exit_code_ok"`
				ImaContainer  bool `json:"ima_container"`
				RPC           bool `json:"rpc"`
				Blocks        bool `json:"blocks"`
			} `json:"healthchecks"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	// Hardware is a struct which holds info of required hardware details for skale node
	Hardware struct {
		Data struct {
			CPUTotalCores       int    `json:"cpu_total_cores"`
			CPUPhysicalCores    int    `json:"cpu_physical_cores"`
			Memory              int64  `json:"memory"`
			Swap                int64  `json:"swap"`
			SystemRelease       string `json:"system_release"`
			UnameVersion        string `json:"uname_version"`
			AttachedStorageSize int64  `json:"attached_storage_size"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	// EndpointStatus holds the information about node endpoint
	EndpointStatus struct {
		Data struct {
			BlockNumber int  `json:"block_number"`
			Syncing     bool `json:"syncing"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	// PublicIPResult is a struct which holds info of public IP
	PublicIPResult struct { //TODO :: have to check
		Data struct {
			PublicIP string `json:"public_ip"`
		} `json:"data"`
	}

	// SslStatus is a struct which holds info if ssl status
	SslStatus struct {
		Data struct {
			IssuedTo       string `json:"issued_to"`
			ExpirationDate string `json:"expiration_date"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	// IMAstatus is a struct which holds info of IMA status
	IMAstatus struct {
		Data []struct {
			SkaleChainName struct {
				Error         string        `json:"error"`
				LastImaErrors []interface{} `json:"last_ima_errors"`
			} `json:"skale-chain-name"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	// BTRFSstatus is a struct which holds btrfs kernal info
	BTRFSstatus struct {
		Data struct {
			KernelModule bool `json:"kernel_module"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	// NodeInfo is a struct which holds Skale Node info
	NodeInfo struct {
		Name           string `json:"name"`
		IP             string `json:"ip"`
		Publicip       string `json:"publicIP"`
		Port           int    `json:"port"`
		StartBlock     int    `json:"start_block"`
		LastRewardDate int    `json:"last_reward_date"`
		FinishTime     int    `json:"finish_time"`
		Status         int    `json:"status"`
		ValidatorID    int    `json:"validator_id"`
		Publickey      string `json:"publicKey"`
		DomainName     string `json:"domain_name"`
		ID             int    `json:"id"`
		Owner          string `json:"owner"`
	}

	// WalletInfo struct holds skale and eth balance
	WalletInfo struct {
		Address         string `json:"address"`
		EthBalanceWei   int64  `json:"eth_balance_wei"`
		SkaleBalanceWei int    `json:"skale_balance_wei"`
		EthBalance      string `json:"eth_balance"`
		SkaleBalance    string `json:"skale_balance"`
	}

	// DBRes holds the ETH and skale balances which stored in DB
	DBRes struct {
		Status string `json:"status"`
		Data   struct {
			Resulttype string `json:"resultType"`
			Result     []struct {
				Metric struct {
					Name                string `json:"__name__"`
					Instance            string `json:"instance"`
					Job                 string `json:"job"`
					SkaleEthBalance     string `json:"skale_eth_balance"`
					SkaleBalance        string `json:"skale_balance"`
					SkaleConAlertCount  string `json:"skale_con_alertCount"`
					SkaleNodeAlertCount string `json:"skale_node_alertCount"`
				} `json:"metric"`
				Value []interface{} `json:"value"`
			} `json:"result"`
		} `json:"data"`
	}

	StatusCore struct {
		Data []struct {
			Image string `json:"image"`
			Name  string `json:"name"`
			State struct {
				Status     string    `json:"Status"`
				Running    bool      `json:"Running"`
				Paused     bool      `json:"Paused"`
				Restarting bool      `json:"Restarting"`
				Oomkilled  bool      `json:"OOMKilled"`
				Dead       bool      `json:"Dead"`
				Pid        int       `json:"Pid"`
				Exitcode   int       `json:"ExitCode"`
				Error      string    `json:"Error"`
				Startedat  time.Time `json:"StartedAt"`
				Finishedat time.Time `json:"FinishedAt"`
				Health     struct {
					Status        string `json:"Status"`
					Failingstreak int    `json:"FailingStreak"`
					Log           []struct {
						Start    time.Time `json:"Start"`
						End      time.Time `json:"End"`
						Exitcode int       `json:"ExitCode"`
						Output   string    `json:"Output"`
					} `json:"Log"`
				} `json:"Health"`
			} `json:"state,omitempty"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}
)
