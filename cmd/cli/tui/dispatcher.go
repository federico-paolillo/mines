package tui

type Subscriber func(intent any)

type Dispatcher struct {
	subscribers []Subscriber
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		subscribers: make([]Subscriber, 0),
	}
}

func (d *Dispatcher) Dispatch(intent any) {
	for _, sub := range d.subscribers {
		sub(intent)
	}
}

func (d *Dispatcher) Subscribe(sub Subscriber) {
	d.subscribers = append(d.subscribers, sub)
}
