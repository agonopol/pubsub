package pubsub

type subscriber struct {
	subs []func(interface{})
}

func Subscriber() *subscriber {
	return &subscriber{make([]func(interface{}),0)}
}

func (this *subscriber) AddSubscriber(sub func(interface{})) {
	this.subs = append(this.subs, sub)
}

func (this *subscriber) Subscribe(pub chan interface{}) {
	for _, sub := range this.subs {
		go func(sub func(interface{})) {
			for v := range pub {
				sub(v)
			}
		}(sub)
	}
}
