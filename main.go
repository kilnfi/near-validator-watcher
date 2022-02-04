package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/skillz-blockchain/near-exporter/api"
	"github.com/skillz-blockchain/near-exporter/collector"
)

func main() {
	rpcAddr := os.Getenv("NEAR_EXPORTER_RPC_ADDR")
	if rpcAddr == "" {
		rpcAddr = "https://rpc.testnet.near.org"
	}

	port := os.Getenv("NEAR_EXPORTER_PORT")
	if port == "" {
		port = "8080"
	}

	trackedAccounts := strings.Split(
		os.Getenv("NEAR_EXPORTER_TRACKED_ACCOUNTS"), ",",
	)

	client := api.NewClient(rpcAddr)

	registry := prometheus.NewPedanticRegistry()
	registry.MustRegister(
		collector.
			NewCollector(client).
			WithTrackedAccounts(trackedAccounts),
	)

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		ErrorLog:      log.New(os.Stderr, log.Prefix(), log.Flags()),
		ErrorHandling: promhttp.ContinueOnError,
	})

	fmt.Println("Listening on port " + port + "...")

	http.Handle("/metrics", handler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
