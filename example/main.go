package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	botgolang "github.com/mail-ru-im/bot-golang"
)

func main() {
	token := os.Getenv("TOKEN")

	bot, err := botgolang.NewBot(token, botgolang.BotDebug(true))
	if err != nil {
		log.Fatalf("cannot connect to bot: %s", err)
	}

	log.Println(bot.Info)

	message := bot.NewTextMessage("d.dorofeev@corp.mail.ru", "Hi")
	if err = message.Send(); err != nil {
		log.Fatalf("failed to send message: %s", err)
	}

	file, err := os.Open("./example.png")
	if err != nil {
		log.Fatalf("cannot open file: %s", err)
	}

	fileMessage := bot.NewFileMessage("d.dorofeev@corp.mail.ru", botgolang.NewMessageFile(file.Name(), file))
	if err := fileMessage.Send(); err != nil {
		log.Println(err)
	}

	if err = fileMessage.Delete(); err != nil {
		log.Fatalf("failed to delete message: %s", err)
	}

	if err = file.Close(); err != nil {
		log.Fatalf("failed to close file: %s", err)
	}

	file, err = os.Open("./voice.aac")
	if err != nil {
		log.Fatalf("cannot open file: %s", err)
	}
	defer file.Close()

	voiceMessage := bot.NewVoiceMessage("g.gabolaev@corp.mail.ru", botgolang.NewMessageFile(file.Name(), file))
	if err := voiceMessage.Send(); err != nil {
		log.Println(err)
	}

	// Simple 30-seconds echo bot with buttons
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	updates := bot.GetUpdatesChannel(ctx)
	for update := range updates {
		fmt.Println(update.Type, update.Payload)
		switch update.Type {
		case botgolang.NEW_MESSAGE:
			message := update.Payload.Message()

			helloBtn := botgolang.NewCallbackButton("Hello", "echo")
			goBtn := botgolang.NewURLButton("go", "https://golang.org/")

			keyboard := botgolang.NewKeyboard()
			keyboard.AddRow(helloBtn, goBtn)

			message.AttachInlineKeyboard(keyboard)

			if err := message.Send(); err != nil {
				log.Printf("failed to send message: %s", err)
			}
		case botgolang.EDITED_MESSAGE:
			message := update.Payload.Message()
			if err := message.Reply("do not edit!"); err != nil {
				log.Printf("failed to reply to message: %s", err)
			}
		case botgolang.CALLBACK_QUERY:
			data := update.Payload.CallbackQuery()
			switch data.CallbackData {
			case "echo":
				response := bot.NewButtonResponse(data.QueryID, "", "Hello World!", false)
				if err := response.Send(); err != nil {
					log.Printf("failed to reply on button click: %s", err)
				}
			}
		}

	}
}
