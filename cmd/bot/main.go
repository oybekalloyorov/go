package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/you/my-telegram-bot/internal/handlers"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN environment variable required")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Bot yaratilmadi: %v", err)
	}
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Graceful shutdown uchun signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		for update := range updates {
			handlers.HandleUpdate(bot, update)
		}
	}()

	<-stop
	log.Println("To'xtatilmoqda...")
	bot.StopReceivingUpdates()
	// ozgina kutamiz — kanallar yopilishi uchun
	time.Sleep(500 * time.Millisecond)
	log.Println("Bot to'xtatildi.")
}
