package main

import (
	"fmt"
	"log"
	"os"

	"github.com/user/ton-price-predictor/internal/config"
	"github.com/user/ton-price-predictor/internal/predictor"
	"github.com/user/ton-price-predictor/internal/price"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fetcher := price.NewFetcher(cfg.APIEndpoint)
	currentPrice, err := fetcher.Fetch()
	if err != nil {
		log.Fatalf("failed to fetch price: %v", err)
	}

	model := predictor.NewModel(cfg.Alpha, cfg.Beta)
	predictedPrice := model.Predict(currentPrice)

	fmt.Printf("Current TON price: $%.4f\n", currentPrice)
	fmt.Printf("Predicted TON price: $%.4f\n", predictedPrice)
	fmt.Printf("Trend: %s\n", model.Trend(currentPrice, predictedPrice))

	os.Exit(0)
}