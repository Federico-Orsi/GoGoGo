package builder

//este Sender sería el Director

type Sender struct{
 sender MessageBuilder
}

func (s *Sender) SetBuilder(mb MessageBuilder){
	s.sender = mb
}

func (s *Sender) BuildMessage(recipient string, message string) (*FormattedMessage, error) {
	s.sender.Set_recipient(recipient)
	s.sender.Set_message(message)
	m, err := s.sender.Build()
	if err != nil {
		return nil, err
	}
    return m, nil
}