package builder

type Message struct {
	Recipient string `json:"recipient" xml:"Recipient"`
	Text      string `json:"text" xml:"textt"`
}

type FormattedMessage struct {
	Body   []byte
	Format string
}