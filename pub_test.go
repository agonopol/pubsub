package pubsub

import (
	"log"
	"testing"
	_"time"
)

func YUNOTestTwoProducers(t *testing.T) {
	pub := Publisher()
	pub.AddPublisher(func() chan interface{} {
		out := make(chan interface{})
		go func() {
			for i := 0; i < 10; i++ {
				out <- i
			}
			close(out)
		}()
		return out
	})
	pub.AddPublisher(func() chan interface{} {
		out := make(chan interface{})
		go func() {
			for i := 10; i > 0; i-- {
				out <- i
			}
			close(out)
		}()
		return out
	})

	consume := make(chan interface{})
	go func() {
		for i := range consume {
			log.Printf("Got %v", i)
		}
	}()
	pub.Publish(consume)
}
