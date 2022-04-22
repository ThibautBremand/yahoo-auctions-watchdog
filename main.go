package main

import (
	"log"
	"os"
	"time"
	"yahoo-auctions-watchdog/cache"
	"yahoo-auctions-watchdog/config"
	"yahoo-auctions-watchdog/coordinator"
	"yahoo-auctions-watchdog/web"
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

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	telegramChatID := os.Getenv("TELEGRAM_CHAT_ID")

	telegramClient := web.NewTelegramClient(telegramToken, telegramChatID)

	c := coordinator.NewCoordinator(
		cfg.Searches,
		sleepPeriod,
		tpl,
		cfg.Changerate,
		telegramClient,
		cfg.DownloadPhotos,
	)
	c.Start(scraped)
}
