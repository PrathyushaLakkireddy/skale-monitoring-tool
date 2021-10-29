package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/config"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/exporter"
	"github.com/PrathyushaLakkireddy/skale-monitoring-tool/monitor"
)

func main() {
	cfg, err := config.ReadFromFile()
	if err != nil {
		log.Fatal(err)
	}

	collector := exporter.NewMetricsCollector(cfg)

	go collector.WatchSlots(cfg)

	go func() {
		for {
			monitor.TelegramAlerting(cfg)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			monitor.ContainerStatusAlert(cfg)
			monitor.NodeStatusAlert(cfg)
			time.Sleep(5 * time.Minute)
		}
	}()

	prometheus.MustRegister(collector)
	http.Handle("/metrics", promhttp.Handler())
	err = http.ListenAndServe(fmt.Sprintf("%s", cfg.Prometheus.ListenAddress), nil)
	if err != nil {
		log.Printf("Error while listening on server : %v", err)
	}
}
