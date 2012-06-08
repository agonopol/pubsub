package pubsub

type publisher struct {
	pubs []func() chan interface{}
	lock chan interface{}
}

func Publisher() *publisher {
	return &publisher{make([]func() chan interface{}, 0), make(chan interface{})}
}

func (this *publisher) AddPublisher(pub func() chan interface{}) {
	this.pubs = append(this.pubs, pub)
}

func (this *publisher) Publish(sub chan interface{}) {
	for _, pub := range this.pubs {
		go func(pub func() chan interface{}) {
			for item := range pub() {
				sub <- item
			}
			this.lock <- pub
		}(pub)
	}
	for _ = range this.pubs {
		_ = <-this.lock
	}
	close(sub)
}
