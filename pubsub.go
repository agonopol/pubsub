package pubsub

type pubSub struct {
	producers []func() chan interface{}
	consumers []func(interface{})
	pipe      chan interface{}
	lock      chan int
}

func New() *pubSub {
	return &pubSub{make([]func() chan interface{}, 0), make([]func(interface{}), 0), make(chan interface{}), make(chan int)}
}

func (this *pubSub) AddProducer(producer func() chan interface{}) {
	this.producers = append(this.producers, producer)
}

func (this *pubSub) AddConsumer(consumer func(interface{})) {
	this.consumers = append(this.consumers, consumer)
}

func (this *pubSub) Run() {
	for _, producer := range this.producers {
		go func(producer func() chan interface{}) {
			for item := range producer() {
				this.pipe <- item
			}
			this.lock <- 0
		}(producer)
	}
	for _,consumer := range this.consumers {
		go func(consumer func(interface{})) {
			for value := range this.pipe {
				consumer(value)
			}

		}(consumer)
	}
	for i := 0; i < len(this.producers); i++ {
		_ = <-this.lock
	}
	close(this.pipe)

}
