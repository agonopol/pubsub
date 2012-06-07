package pubsub

import (
	"log"
	"testing"
	"time"
)

func TestOneProducerTwoConsumers(t *testing.T) {
	pub := New()
	pub.AddProducer(func() chan interface{} {
		out := make(chan interface{})
		go func() {
			for i := 0; i < 10; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	})
	pub.AddConsumer(func(x interface{}) {
		log.Printf("Consumer [%v] consumed [%v]", 0, x)
		time.Sleep(time.Second)
	})
	pub.AddConsumer(func(x interface{}) {
		log.Printf("Consumer [%v] consumed [%v]", 1, x)
		time.Sleep(time.Second)
	})

	pub.Run()
}
