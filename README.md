# Golang interface for ICQ bot API

## Install
```bash
go get github.com/DmitryDorofeev/goicqbot
```

## Usage

Create your own bot by sending the /newbot command to Metabot and follow the instructions.

Note a bot can only reply after the user has added it to his contact list, or if the user was the first to start a dialogue.

### Create your bot

```go
package main

import "github.com/DmitryDorofeev/goicqbot"

func main() {
    bot := goicqbot.NewBot(BOT_TOKEN)

    bot.sendMessage("chat_id", "Hello")
}
```

### Send message

```go
bot.SendMessage("mychat@agent.chat", "text")
```

### Passing options

You can override bot's API URL:

```go
bot := goicqbot.NewBot(BOT_TOKEN, goicqbot.BotApiUrl("https://agent.mail.ru/bot/v1"))
```



## Roadmap

- [x] Send message

- [ ] Events subscription

- [ ] Send files

- [ ] Godoc

- [ ] Tests

- [ ] Send voice

- [ ] Delete message

- [ ] Edit message