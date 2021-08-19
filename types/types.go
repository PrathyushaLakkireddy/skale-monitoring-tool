package types

import (
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

	// Balance struct which holds information of Account Balancce
	Balance struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Context struct {
				Slot int `json:"slot"`
			} `json:"context"`
			Value int64 `json:"value"`
		} `json:"result"`
		ID int `json:"id"`
	}

	// AccountInfo struct which holds Account Information
	AccountInfo struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Context struct {
				Slot int `json:"slot"`
			} `json:"context"`
			Value struct {
				Data struct {
					Nonce struct {
						Initialized struct {
							Authority     string `json:"authority"`
							Blockhash     string `json:"blockhash"`
							FeeCalculator struct {
								LamportsPerSignature int `json:"lamportsPerSignature"`
							} `json:"feeCalculator"`
						} `json:"initialized"`
					} `json:"nonce"`
				} `json:"data"`
				Executable bool   `json:"executable"`
				Lamports   int    `json:"lamports"`
				Owner      string `json:"owner"`
				RentEpoch  int    `json:"rentEpoch"`
			} `json:"value"`
		} `json:"result"`
		ID int `json:"id"`
	}

	// EpochInfo struct which holds information of current Epoch
	EpochInfo struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			AbsoluteSlot int64 `json:"absoluteSlot"`
			BlockHeight  int64 `json:"blockHeight"`
			Epoch        int64 `json:"epoch"`
			SlotIndex    int64 `json:"slotIndex"`
			SlotsInEpoch int64 `json:"slotsInEpoch"`
		} `json:"result"`
		ID int `json:"id"`
	}

	// EpochShedule struct holds Epoch Shedule Information
	EpochShedule struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			FirstNormalEpoch         int  `json:"firstNormalEpoch"`
			FirstNormalSlot          int  `json:"firstNormalSlot"`
			LeaderScheduleSlotOffset int  `json:"leaderScheduleSlotOffset"`
			SlotsPerEpoch            int  `json:"slotsPerEpoch"`
			Warmup                   bool `json:"warmup"`
		} `json:"result"`
		ID int `json:"id"`
	}

	// LeaderShedule struct holds information of leader schedule for an epoch
	LeaderShedule struct {
		Jsonrpc string             `json:"jsonrpc"`
		Result  map[string][]int64 `json:"result"`
		ID      int                `json:"id"`
	}

	// ConfirmedBlocks struct which holds information of confirmed blocks
	ConfirmedBlocks struct {
		Jsonrpc string  `json:"jsonrpc"`
		Result  []int64 `json:"result"`
		ID      int     `json:"id"`
	}

	// BlockTime struct which holds information of estimated production time of a confirmed block
	BlockTime struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  int64  `json:"result"`
	}

	// NodeHealth struct which holds information of health of the node
	NodeHealth struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  string `json:"result"`
		Error   struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
			Data    struct {
			} `json:"data"`
		} `json:"error"`
	}

	// Version struct which holds information of solana version
	EthResult struct {
		// Jsonrpc string `json:"jsonrpc"`
		// Result struct {
		// 	SolanaCore string `json:"solana-core"`
		// } `json:"result"`

		// type AutoGenerated struct {
		// 	ID      int    `json:"id"`
		Jsonrpc string `json:"jsonrpc"`
		Result  string `json:"result"`
		// }
	}

	// Identity struct holds the pubkey for the current node
	Identity struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			Identity string `json:"identity"`
		} `json:"result"`
	}

	// VoteAccount struct holds information of vote account
	VoteAccount struct {
		ActivatedStake   int64     `json:"activatedStake"`
		Commission       int64     `json:"commission"`
		EpochCredits     [][]int64 `json:"epochCredits"`
		EpochVoteAccount bool      `json:"epochVoteAccount"`
		LastVote         int       `json:"lastVote"`
		NodePubkey       string    `json:"nodePubkey"`
		RootSlot         int       `json:"rootSlot"`
		VotePubkey       string    `json:"votePubkey"`
	}

	// rpcError struct which holds Error message of RPC
	rpcError struct {
		Message string `json:"message"`
		Code    int64  `json:"id"`
	}

	// ConfirmedBlock struct which holds blocktime of confirmedBlock at current slot height
	ConfirmedBlock struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			BlockTime int64 `json:"blockTime"`
		} `json:"result"`
	}

	Syncing struct {
		// ID      int    `json:"id"`
		Jsonrpc string `json:"jsonrpc"`
		Result  bool   `json:"result"`
	}

	StatusCore struct {
		Data []struct {
			Image string `json:"image"`
			Name  string `json:"name"`
			State struct {
				Status     string `json:"Status"`
				Running    bool   `json:"Running"`
				Paused     bool   `json:"Paused"`
				Restarting bool   `json:"Restarting"`
				OOMKilled  bool   `json:"OOMKilled"`
				Dead       bool   `json:"Dead"`
				Pid        int    `json:"Pid"`
				ExitCode   int    `json:"ExitCode"`
				Error      string `json:"Error"`
				StartedAt  string `json:"StartedAt"`
				FinishedAt string `json:"FinishedAt"`
			} `json:"state,omitempty"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	// SGXStatus server info - connection status and SGX wallet version
	SGXStatus struct {
		Data struct {
			Status           int    `json:"status"`
			StatusName       string `json:"status_name"`
			SgxWalletVersion string `json:"sgx_wallet_version"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

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

	EndpointStatus struct {
		Data struct {
			BlockNumber int  `json:"block_number"`
			Syncing     bool `json:"syncing"`
		} `json:"data"`
		Error interface{} `json:"error"`
	}

	PublicIPResult struct { //TODO :: have to check
		Data struct {
			IP string `json:"ip"`
		} `json:"data"`
	}
)