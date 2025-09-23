package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleUpdate - yagona joydan barcha updatelarni qayta ishlash uchun oddiy funksiya.
// Bu yerga o'zingizning business logic, DB chaqiriqlari va hokazolarni qo'shing.
func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Matnli xabar
	if update.Message != nil {
		chatID := update.Message.Chat.ID

		// komandalarni ishlash (masalan /start, /help)
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(chatID, "Assalomu alaykum! Men Go bilan yozilgan botman.")
				bot.Send(msg)
			case "keyboard":
				keyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Ha", "choice_yes"),
						tgbotapi.NewInlineKeyboardButtonData("Yo'q", "choice_no"),
					),
				)
				msg := tgbotapi.NewMessage(chatID, "Tanlang:")
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(chatID, "Noma'lum buyruq.")
				bot.Send(msg)
			}
			return
		}

		// oddiy echo — demo maqsadida
		if update.Message.Text != "" {
			msg := tgbotapi.NewMessage(chatID, "Siz yozdingiz: "+update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Send error: %v", err)
			}
		}
		return
	}

	// Callback query (inline keyboard tugmachalari)
	if update.CallbackQuery != nil {
		cb := update.CallbackQuery

		// javob beramiz (spinner ni o'chirish uchun)
		if _, err := bot.Request(tgbotapi.NewCallback(cb.ID, "Qabul qilindi")); err != nil {
			log.Printf("Callback answer error: %v", err)
		}

		// chatga xabar ham yuborish mumkin
		if _, err := bot.Send(tgbotapi.NewMessage(cb.Message.Chat.ID, "Siz tanladingiz: "+cb.Data)); err != nil {
			log.Printf("Send message after callback error: %v", err)
		}
		return
	}
}
