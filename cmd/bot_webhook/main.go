package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/you/my-telegram-bot/internal/handlers"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	webhookURL := os.Getenv("q") // masalan https://abcd-1234.ngrok.io

	if token == "" || webhookURL == "" {
		log.Fatal("TELEGRAM_TOKEN va WEBHOOK_URL muhit o'zgaruvchilari kerak")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Bot yaratilmadi: %v", err)
	}
	bot.Debug = false

	// ✅ To‘g‘rilangan joy
	wh, err := tgbotapi.NewWebhook(webhookURL + "/" + bot.Token)
	if err != nil {
		log.Fatalf("Webhook yaratishda xato: %v", err)
	}

	// Telegramga webhook sozlash
	_, err = bot.Request(wh)
	if err != nil {
		log.Fatalf("Webhook qo'yishda xato: %v", err)
	}

	info, _ := bot.GetWebhookInfo()
	if info.LastErrorDate != 0 {
		log.Printf("Webhook info error: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)

	// Local HTTP server
	go func() {
		log.Println("Serverni 0.0.0.0:8080 da ishga tushiryapman")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	log.Printf("Listening webhook at %s/%s", webhookURL, bot.Token)
	for update := range updates {
		handlers.HandleUpdate(bot, update)
	}
}
