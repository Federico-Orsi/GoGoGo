package builder

import "encoding/json"

type JSON_message struct {
	message Message
}

func (m *JSON_message) Set_recipient(recipient string) MessageBuilder {
	m.message.Recipient = recipient
	return m
}

func (m *JSON_message) Set_message(text string) MessageBuilder {
	m.message.Text = text
	return m
}

func (m *JSON_message) Build() (*FormattedMessage, error) {
	data, err := json.Marshal(m.message)

	if err != nil {
       return nil, err
	}

	return &FormattedMessage{Body: data, Format: "json"}, nil
}