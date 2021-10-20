package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// Telegram bot details struct
	Telegram struct {
		// BotToken is the token of your telegram bot
		BotToken string `mapstructure:"tg_bot_token"`
		// ChatID is the id of telegarm chat which will be used to get alerts
		ChatID int64 `mapstructure:"tg_chat_id"`
	}

	// SendGrid stores sendgrid API credentials
	SendGrid struct {
		// Token of sendgrid account
		Token string `mapstructure:"sendgrid_token"`
		// ToEmailAddress is the email to which all the alerts will be sent
		ReceiverEmailAddress string `mapstructure:"receiver_email_address"`
		// SendgridEmail is the email of sendgrid account which will be used to send mail alerts
		SendgridEmail string `mapstructure:"account_email"`
		// SendgridName is the name of sendgrid account which will be used to send mail alerts
		SendgridName string `mapstructure:"sendgrid_account_name"`
	}

	// Scraper defines the time intervals for multiple scrapers to fetch the data
	Scraper struct {
		// Rate is to call and get the data for specified targets on that particular time interval
		Rate string `mapstructure:"rate"`
	}

	// Prometheus stores Prometheus details
	Prometheus struct {
		// ListenAddress to export metrics on the given port
		ListenAddress string `mapstructure:"listen_address"`
		// PrometheusAddress to connect to prormetheus where it has running
		PrometheusAddress string `mapstructure:"prometheus_address"`
	}

	// Endpoints defines multiple API base-urls to fetch the data
	Endpoints struct {
		// RPCEndPoint is used to gather information about validator sgx wallet status, block number, core status etc ...
		SkaleNodeIP string `mapstructure:"skale_node_ip"`
	}

	// EnableAlerts struct which holds options to enalbe/disable alerts
	EnableAlerts struct {
		// EnableTelegramAlerts which takes an option to enable/disable telegram alerts
		EnableTelegramAlerts bool `mapstructure:"enable_telegram_alerts"`
		// EnableTelegramAlerts which takes an option to enable/disable emial alerts
		EnableEmailAlerts bool `mapstructure:"enable_email_alerts"`
	}

	// RegularStatusAlerts defines time-slots to receive validator status alerts
	RegularStatusAlerts struct {
		// AlertTimings is the array of time slots to send validator status alerts at that particular timings
		AlertTimings []string `mapstructure:"alert_timings"`
	}

	// AlerterPreferences which holds individual alert settings which takes an option to  enable/disable particular alert
	AlerterPreferences struct {
		// NodeHealthAlert which takes an option to  enable/disable node Health status alert, on enable sends alerts
		NodeHealthAlert string `mapstructure:"node_health_alert"`
		// NodeStatusAlert            string `mapstructure:"node_status_alert"`
		// BlockSyncAlerts which takes an option to disable/enable Block synching status alerts, on enable send alert
		// when block is in synching process
		BlockSyncAlerts string `mapstructure:"block_sync_alerts"`
		// ContainerHealthAlerts which takes an option to disable/enable Container heath status alerts, on enable send alert
		// when container stopped running or paused or dead
		ContainerHealthAlerts string `mapstructure:"container_health_alerts"`
		// EthbalanceChangeAlerts which takes an option to disable/enable Account balance change alerts, on enable sends alert
		// when balance has dropped to eth balance threshold
		EthbalanceChangeAlerts string `mapstructure:"ethbalance_change_alerts"`
		// SklbalanceChangeAlerts which takes an option to disable/enable Account balance change alerts, on enable sends alert
		// when balance has dropped to skale balance threshold
		SklbalanceChangeAlerts string `mapstructure:"sklbalance_change_alerts"`
		// ETHDelegationAlerts which takes an option to disable/enable Account balance delegation alerts, on enable sends alerts
		// when ETH balance changes
		ETHDelegationAlerts string `mapstructure:"eth_delegation_alerts"`
		// SKLDelegationAlerts which takes an option to disable/enable Account balance delegation alerts, on enable sends alerts
		// when SKALE balance changes
		SKLDelegationAlerts string `mapstructure:"skl_delegation_alerts"`

		BTRFSstatusAlerts string `mapstructure:"btrfs_status_alerts"`

		SGXstatusAlerts string `mapstructure:"sgx_status_alerts"`
	}

	//  AlertingThreshold defines threshold condition for different alert-cases.
	//`Alerter` will send alerts if the condition reaches the threshold
	AlertingThreshold struct {
		EthbalanceThreshold float64 `mapstructure:"ethbalance_change_threshold"`
		SklbalanceThreshold float64 `mapstructure:"sklbalance_change_threshold"`
	}

	// Config defines all the configurations required for the app
	Config struct {
		Endpoints           Endpoints           `mapstructure:"skale_endpoint"`
		EnableAlerts        EnableAlerts        `mapstructure:"enable_alerts"`
		RegularStatusAlerts RegularStatusAlerts `mapstructure:"regular_status_alerts"`
		AlerterPreferences  AlerterPreferences  `mapstructure:"alerter_preferences"`
		AlertingThresholds  AlertingThreshold   `mapstructure:"alerting_threholds"`
		Scraper             Scraper             `mapstructure:"scraper"`
		Telegram            Telegram            `mapstructure:"telegram"`
		SendGrid            SendGrid            `mapstructure:"sendgrid"`
		Prometheus          Prometheus          `mapstructure:"prometheus"`
	}
)

// ReadFromFile to read config details using viper
func ReadFromFile() (*Config, error) {
	// usr, err := user.Current()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// configPath := path.Join(usr.HomeDir, `.solana-tool/config/`)
	// log.Printf("Config Path : %s", configPath)

	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	// v.AddConfigPath(configPath)
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config.toml: %v", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshaling config.toml to application config: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("error occurred in config validation: %v", err)
	}

	return &cfg, nil
}

// Validate config struct
func (c *Config) Validate(e ...string) error {
	v := validator.New()
	if len(e) == 0 {
		return v.Struct(c)
	}
	return v.StructExcept(c, e...)
}
