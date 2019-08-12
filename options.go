package goicqbot

type BotOption interface {
	Type() string
	String() string
	Int() int
	Bool() bool
}

type BotApiUrl string

func (o BotApiUrl) Type() string {
	return "api_url"
}

func (o BotApiUrl) String() string {
	return string(o)
}

func (o BotApiUrl) Int() int {
	return 0
}

func (o BotApiUrl) Bool() bool {
	return false
}

type BotDebug bool

func (o BotDebug) Type() string {
	return "debug"
}

func (o BotDebug) String() string {
	return ""
}

func (o BotDebug) Int() int {
	return 0
}

func (o BotDebug) Bool() bool {
	return bool(o)
}