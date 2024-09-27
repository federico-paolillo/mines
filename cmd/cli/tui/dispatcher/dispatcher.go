package dispatcher

type Subscriber = func(intent any)
type Unsubscribe = func()

type Dispatcher struct {
	subscribers      map[int]Subscriber
	subscribersIndex int
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		subscribers: make(map[int]Subscriber, 0),
	}
}

func (d *Dispatcher) Dispatch(intent any) {
	for _, sub := range d.subscribers {
		sub(intent)
	}
}

func (d *Dispatcher) Subscribe(sub Subscriber) Unsubscribe {
	currentIndex := d.subscribersIndex

	d.subscribers[currentIndex] = sub

	d.subscribersIndex++

	return func() {
		delete(d.subscribers, currentIndex)
	}
}
