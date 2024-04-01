package builder

import "encoding/xml"

type XML_message struct {
	message Message
}

func (m *XML_message) Set_recipient(recipient string) MessageBuilder {
	m.message.Recipient = recipient
	return m
}

func (m *XML_message) Set_message(text string) MessageBuilder {
	m.message.Text = text
	return m
}

func (m *XML_message) Build() (*FormattedMessage, error) {
	data, err := xml.Marshal(m.message)

	if err != nil {
       return nil, err
	}

	return &FormattedMessage{Body: data, Format: "xml"}, nil
}