<img src="logo_bot.png" width="100" height="100">

# VK Teams Bot API for Golang
[![Go](https://github.com/mail-ru-im/bot-golang/actions/workflows/go.yml/badge.svg)](https://github.com/mail-ru-im/bot-golang/actions/workflows/go.yml)
[![codecov](https://codecov.io/github/mail-ru-im/bot-golang/graph/badge.svg?token=0HX8DY24SR)](https://codecov.io/github/mail-ru-im/bot-golang)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/mail-ru-im/bot-golang)
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)

### [<img src="logo_msg.png" width="16"> VK Teams API Specification](https://teams.vk.com/botapi/)

## Getting started

* Create your own bot by sending the _/newbot_ command to _Metabot_ and follow the instructions.
    >Note: a bot can only reply after the user has added it to his contact list, or if the user was the first to start a dialogue.
* You can configure the domain that hosts your VK Teams server. When instantiating the Bot class, add the address of your domain.
* An example of how to use the framework can be seen in _example/main.go_

## Install
```bash
go get github.com/mail-ru-im/bot-golang
```

## Usage

Create your own bot by sending the /newbot command to _Metabot_ and follow the instructions.
Note a bot can only reply after the user has added it to his contacts list, or if the user was the first to start a dialogue.

### Create your bot

```go
package main

import "github.com/mail-ru-im/bot-golang"

func main() {
    bot, err := botgolang.NewBot(BOT_TOKEN)
    if err != nil {
        log.Println("wrong token")
    }

    message := bot.NewTextMessage("some@mail.com", "text")
    message.Send()
}
```

### Send and edit messages

You can create, edit and reply to messages like a piece of cake.

```go
message := bot.NewTextMessage("some@mail.com", "text")
message.Send()

message.Text = "new text"

message.Edit()
message.Reply("I changed my text")
```

### Subscribe events

Get all updates from the channel. Use context for cancellation.

```go
ctx, finish := context.WithCancel(context.Background())
updates := bot.GetUpdatesChannel(ctx)
for update := range updates {
	// your logic here
}
```

### Passing options

You don't need this.
But if you do, you can override bot's API URL:

```go
bot := botgolang.NewBot(BOT_TOKEN, botgolang.BotApiURL("https://vkteams.com/bot/v1"))
```
And debug all api requests and responses:

```go
bot := botgolang.NewBot(BOT_TOKEN, botgolang.BotDebug(true))
```
