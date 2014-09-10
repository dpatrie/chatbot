package bot

type PluginEcho struct{}

func (p PluginEcho) Name() string {
	return "Echo v1.0"
}

func (p PluginEcho) Execute(msg Message, bot Bot) error {
	if msg.From() != bot.Name() {
		bot.Send(msg.Body())
	}
	return nil
}
