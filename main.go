package main

import (
	"log"
	"time"
	"yahoo-auctions-watchdog/cache"
	"yahoo-auctions-watchdog/config"
	"yahoo-auctions-watchdog/coordinator"
)

func main() {
	log.Println("Starting yahoo-auctions-watchdog")

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error while loading config: %v", err)
	}

	tpl, err := cfg.LoadTemplate()
	if err != nil {
		log.Fatalf("Could not parse message template: %v\n", err)
	}

	scraped, err := cache.LoadCache()
	if err != nil {
		log.Fatalf("Could not load cache: %v", err)
	}

	sleepPeriod := time.Duration(cfg.Delay) * time.Second

	c := coordinator.NewCoordinator(cfg.Searches, sleepPeriod, tpl, cfg.Changerate)
	c.Start(scraped)
}
