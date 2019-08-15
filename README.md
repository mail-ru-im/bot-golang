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

Note a bot can only reply after the user has added it to his contacts list, or if the user was the first to start a dialogue.

### Create your bot

```go
package main

import "github.com/DmitryDorofeev/goicqbot"

func main() {
    bot, err := goicqbot.NewBot(BOT_TOKEN)
    if err != nil {
        log.Println("wrong token")
    }

    message := bot.NewTextMessage("awesomechat@agent.chat", "text")
    message.Send()
}
```

### Send and edit messages

You can create, edit and reply to messages like a piece of cake.

```go
message := bot.NewTextMessage("awesomechat@agent.chat", "text")
message.Send()

fmt.Println(message.MsgID)

message.Text = "new text"

message.Edit()
// AWESOME!

message.Reply("hey, what did you write before???")
// SO HOT!
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

- [x] Send files

- [x] Delete message

- [ ] Send voice

