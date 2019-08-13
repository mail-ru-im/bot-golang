<img src="https://github.com/mail-ru-im/bot-python/blob/master/logo.png" width="100" height="100">

# Golang interface for ICQ bot API
[![CircleCI](https://circleci.com/gh/DmitryDorofeev/goicqbot.svg?style=svg)](https://circleci.com/gh/DmitryDorofeev/goicqbot)

 - *Brand new Bot API!*

 - *Zero-configuration library*

 - *Simple and clear interface*

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

    bot.SendMessage(&goicqbot.Message{Text: "text", ChatID: "awesomechat@agent.chat"})
}
```

### Send and edit message

You can create and edit messages like a piece of cake.

```go
message := &goicqbot.Message{Text: "text", ChatID: "awesomechat@agent.chat"}
bot.SendMessage(message)

fmt.Println(message.MsgID)

message.Text = "new text"

bot.EditMessage(message)
// AWESOME!
```

### Subscribe events

Get all updates from the channel. Use context for cancellation.

```go
ctx, finish := context.WithCancel(context.Background())
updates := bot.GetUpdatesChannel(ctx)
for update := range updates {
	// your awesome logic here
}
```

### Passing options

You don't need this.
But if you do, you can override bot's API URL:

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

- [x] Tests

- [x] Godoc

- [x] Edit message

- [ ] Send files

- [ ] Send voice

- [ ] Delete message
