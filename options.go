package goicqbot

type BotOption interface {
	Type() string
	Value() interface{}
}

type BotApiUrl string

func (o BotApiUrl) Type() string {
	return "api_url"
}

func (o BotApiUrl) Value() interface{} {
	return string(o)
}

type BotDebug bool

func (o BotDebug) Type() string {
	return "debug"
}

func (o BotDebug) Value() interface{} {
	return bool(o)
}
