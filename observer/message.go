package observer

type Message struct {
	observers map[string]Observer
	Msg       string
}

func (m *Message) Add_Observer(name string, o Observer) {
	if m.observers == nil {
		m.observers = make(map[string]Observer) //make es un metodo propio de los maps para crearlos!!
	}

	m.observers[name] = o
}

func (m *Message) Remove_Observer(name string) {
	delete(m.observers, name) // delete() es un metodo propio de los maps!!
}

func (m *Message) Notify_Observers() {

	for _, values := range m.observers {

		values.Notify(m.Msg)
	}
}
