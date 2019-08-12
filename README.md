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

    bot.sendMessage(goicqbot.Message{Text: "text", ChatID: "awesomechat@agent.chat"})
}
```

### Send message

```go
bot.SendMessage(goicqbot.Message{Text: "text", ChatID: "awesomechat@agent.chat"})
```

### Subscribe events

```go
updates := bot.GetUpdatesChannel()

for update := range updates {
	// your awesome logic here
}
```

### Passing options

You can override bot's API URL:

```go
bot := goicqbot.NewBot(BOT_TOKEN, goicqbot.BotApiUrl("https://agent.mail.ru/bot/v1"))
```

And debug all api requests and responses:

```go
bot := goicqbot.NewBot(BOT_TOKEN, goicqbot.BotDebug(true))
```


## Roadmap

- [x] Send message

- [x] Events subscription

- [ ] Send files

- [ ] Godoc

- [ ] Tests

- [ ] Send voice

- [ ] Delete message

- [ ] Edit message