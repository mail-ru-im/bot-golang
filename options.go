package goicqbot

type BotOption interface {
	Type() string
	String() string
	Int() int
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