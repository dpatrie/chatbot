package bot

type Bot interface {
	Name() string
	Send(msg string)
	Connect() error
	Listen(recv chan Message)
}

//*************************************************

type Message interface {
	Body() string
	From() string
}

//*************************************************

type Plugin interface {
	Name() string
	Execute(msg Message, bot Bot) error
}
