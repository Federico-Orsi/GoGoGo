package builder

type MessageBuilder interface {
	Set_recipient(string) MessageBuilder
	Set_message(string) MessageBuilder
	Build() (*FormattedMessage, error)
}