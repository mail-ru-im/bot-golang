package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mail-ru-im/bot-golang"
)

func main() {
	token := os.Getenv("TOKEN")

	bot, err := botgolang.NewBot(token, botgolang.BotDebug(true))
	if err != nil {
		log.Fatalf("cannot connect to bot: %s", err)
	}

	log.Println(bot.Info)

	message := bot.NewTextMessage("d.dorofeev@corp.mail.ru", "Hi")
	message.Send()

	file, err := os.Open("./example.png")
	if err != nil {
		log.Fatalf("cannot open file: %s", err)
	}

	fileMessage := bot.NewFileMessage("d.dorofeev@corp.mail.ru", file)
	if err := fileMessage.Send(); err != nil {
		log.Println(err)
	}

	fileMessage.Delete()
	file.Close()

	file, err = os.Open("./voice.aac")
	if err != nil {
		log.Fatalf("cannot open file: %s", err)
	}
	defer file.Close()

	voiceMessage := bot.NewVoiceMessage("g.gabolaev@corp.mail.ru", file)
	if err := voiceMessage.Send(); err != nil {
		log.Println(err)
	}

	// Simple 30-seconds echo bot
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	updates := bot.GetUpdatesChannel(ctx)
	for update := range updates {
		fmt.Println(update.Type, update.Payload)
		switch update.Type {
		case botgolang.NEW_MESSAGE:
			message := update.Payload.Message()
			if err := message.Send(); err != nil {
				log.Printf("something went wrong: %s", err)
			}
		case botgolang.EDITED_MESSAGE:
			message := update.Payload.Message()
			message.Reply("do not edit!")
		}

	}
}
