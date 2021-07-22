package main

import (
	"context"
	botgolang "github.com/mail-ru-im/bot-golang"
	"log"
	"time"
)


func main() {
	token := "YOUR TOKEN HERE :)"

	bot, err := botgolang.NewBot(token, botgolang.BotApiURL("https://api.internal.myteam.mail.ru/bot/v1"), botgolang.BotDebug(true))
	if err != nil {
		log.Fatalf("cannot connect to bot: %s", err)
	}

	log.Println(bot.Info)

	message := bot.NewTextMessage("d.udovichenko@corp.mail.ru", "Send me anything")
	if err = message.Send(); err != nil {
		log.Fatalf("failed to send message: %s", err)
	}

	// Simple 30-seconds echo bot with buttons
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	updates := bot.GetUpdatesChannel(ctx)
	for update := range updates {
		switch update.Type {
		case botgolang.NEW_MESSAGE:
			message := update.Payload.Message()

			helloBtn := botgolang.NewCallbackButton("Click me to edit this message", "nevermind")
			keyboard := botgolang.NewKeyboard()
			keyboard.AddRow(helloBtn)

			message.AttachInlineKeyboard(keyboard)

			if err := message.Send(); err != nil {
				log.Printf("failed to send message: %s", err)
			}
		case botgolang.CALLBACK_QUERY:

			editedMsg := bot.NewMessage(update.Payload.Msg.Chat.ID)
			editedMsg.ID = update.Payload.Msg.MsgID
			editedMsg.Text =  "This message was edited"

			if err := editedMsg.Edit(); err != nil {
				log.Fatalf("failed to edit message: %s", err)
			}
		}

	}
}

