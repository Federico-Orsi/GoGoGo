package observer

type Observable interface {
	Add_Observer(name string, observer Observer)
	Remove_Observer(name string)
	Notify_Observers()
}