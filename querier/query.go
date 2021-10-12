package querier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/types"
)

// GetETHbalanceFromDB returns ETH balance from prometheus
func GetETHbalanceFromDB(cfg *config.Config) (string, error) {
	var result types.DBRes
	var bal string
	response, err := http.Get(fmt.Sprintf("%s/api/v1/query=skale_eth_balance", cfg.Prometheus.PrometheusAddress))
	if err != nil {
		log.Printf("Error while getting ETH balance from DB : %v", err)
		return bal, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		log.Printf("Error while unmarshalling accont balance : %v", err)
	}
	if len(result.Data.Result) > 0 {
		bal = result.Data.Result[0].Metric.SkaleEthBalance
	}
	log.Printf("ETH bal from db: %v", bal)

	return bal, nil
}

// GetSKLbalanceFromDB returns skale balance from prometheus
func GetSKLbalanceFromDB(cfg *config.Config) (string, error) {
	var result types.DBRes
	var bal string
	response, err := http.Get(fmt.Sprintf("%s/api/v1/query=skale_balance", cfg.Prometheus.PrometheusAddress))
	if err != nil {
		log.Printf("Error while getting skale balance from DB : %v", err)
		return bal, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(responseData, &result)
	if err != nil {
		log.Printf("Error while unmarshalling accont balance : %v", err)
	}
	if len(result.Data.Result) > 0 {
		bal = result.Data.Result[0].Metric.SkaleBalance
	}
	log.Printf("skale bal from db: %v", bal)

	return bal, nil
}
