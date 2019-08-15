package main

import (
	"github.com/DmitryDorofeev/goicqbot"
	"log"
	"os"
)

func main() {
	token := os.Getenv("TOKEN")

	bot, err := goicqbot.NewBot(token, goicqbot.BotDebug(true))
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
}
