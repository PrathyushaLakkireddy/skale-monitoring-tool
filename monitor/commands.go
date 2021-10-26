package monitor

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// TelegramAlerting will check for the commands from the configured telegram account
// If any commands are given in the tg account then Alerter will send the response back according to the input
func TelegramAlerting(cfg *config.Config) {
	if strings.ToUpper(strconv.FormatBool(cfg.EnableAlerts.EnableTelegramAlerts)) == "FALSE" {
		return
	}
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.BotToken)
	if err != nil {
		log.Fatalf("Please configure telegram bot token %v:", err)
		return
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	msgToSend := ""

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/eth_balance" {
			msgToSend = GetETHBal(cfg)
		} else if update.Message.Text == "/skale_balance" {
			msgToSend = GetSklBal(cfg)
		} else if update.Message.Text == "/block_number" {
			msgToSend = GetBlockNum(cfg)
		} else if update.Message.Text == "/sgx_status" {
			msgToSend = GetSGXstat(cfg)
		} else if update.Message.Text == "/container_status" {
			msgToSend = GetContainerHealth(cfg)
		} else if update.Message.Text == "/stop" {
			msgToSend = Stop()
			if msgToSend != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgToSend)
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
			log.Fatalf(msgToSend)
		} else if update.Message.Text == "/list" {
			msgToSend = GetHelp()
		} else {
			text := strings.Split(update.Message.Text, "")
			if len(text) != 0 {
				if text[0] == "/" {
					msgToSend = "Command not found do /list to know about available commands"
				} else {
					msgToSend = " "
				}
			}
		}

		log.Printf("[%s] %s", update.Message.From.UserName, msgToSend)

		if msgToSend != " " {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgToSend)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}

	}

}

// GetHelp returns the msg to show for /help
func GetHelp() string {
	msg := "List of available commands\n /eth_balance - returns validator ETH balance\n /skale_balance - return validator SKL balance\n" +
		"/block_number - returns current block number\n /sgx_status - returns sgx server status\n /container_status - returns status of the containers\n /stop - which panics the running code and also alerts will be stopped\n /list - list out the available commands"
	return msg
}

// GetETHBal returns ETH balance
func GetETHBal(cfg *config.Config) string {
	var msg string

	res, err := GetWalletInfo(cfg)
	if err != nil {
		log.Printf("Error while getting ETH balance : %v", err)
	}
	bal, _ := strconv.ParseFloat(res.EthBalance, 64)
	b := fmt.Sprintf("%.2f", bal)
	msg = fmt.Sprintf("Your ETH balance is %sETH\n", b)
	return msg
}

// GetSklBal returns skale balance
func GetSklBal(cfg *config.Config) string {
	var msg string

	res, err := GetWalletInfo(cfg)
	if err != nil {
		log.Printf("Error while getting skale balance : %v", err)
	}
	bal, _ := strconv.ParseFloat(res.SkaleBalance, 64)
	b := fmt.Sprintf("%.2f", bal)
	msg = fmt.Sprintf("Your SKALE balance is %sSKL\n", b)
	return msg
}

// GetBlockNum returns current block number
func GetBlockNum(cfg *config.Config) string {
	var msg string

	res, err := GetEndpointStatus(cfg)
	if err != nil {
		log.Printf("Error while getting blocknumber : %v", err)
	}
	msg = msg + fmt.Sprintf("Your current block number is : %d\n", res.Data.BlockNumber)
	return msg
}

// GetSGXstat returns sgx status
func GetSGXstat(cfg *config.Config) string {
	var msg string

	res, err := GetSGXStatus(cfg)
	if err != nil {
		log.Printf("Error while getting SGX status : %v", err)
	}
	s := res.Data.StatusName
	msg = msg + fmt.Sprintf("Yout SGX server is %s\n", s)
	return msg
}

// GetContainerHealth returns containers health status
func GetContainerHealth(cfg *config.Config) string {
	var msg string

	res, err := GetCoreStatus(cfg)
	if err != nil {
		log.Printf("Error while getting containers status : %v", err)
	}
	for i, container := range res.Data {
		if res.Data[i].State.Health.Status != "" {
			msg = msg + fmt.Sprintf("%s container is %s\n", container.Name, res.Data[i].State.Health.Status)
		} else {
			if res.Data[i].State.Running == true {
				msg = msg + fmt.Sprintf("%s container is Running\n", container.Name)
			} else if res.Data[i].State.Running == false {
				msg = msg + fmt.Sprintf("%s container is not Running \n", container.Name)
			} else if res.Data[i].State.Paused == true {
				msg = msg + fmt.Sprintf("%s container is Paused \n", container.Name)
			} else if res.Data[i].State.Dead == true {
				msg = msg + fmt.Sprintf("%s container is Dead \n", container.Name)
			}
		}

	}

	return msg
}

// Stop which will be used to stop the the running program of monitoring tool
func Stop() string {
	var msg string
	msg = "Monitoring tool has stopped"
	return msg
}
