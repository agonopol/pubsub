package pubsub

import (
	"log"
	"testing"
	"time"
)

func TestTwoSubscribers(t *testing.T) {
	sub := Subscriber()
	sub.AddSubscriber(func(i interface{}) {
		log.Printf("First subscriber consumed %v",i)
		time.Sleep(time.Millisecond)
	})
	sub.AddSubscriber(func(i interface{}) {
		log.Printf("Second subscriber consumed %v",i)
		time.Sleep(time.Millisecond)
	})
	pipe := make(chan interface{})
	sub.Subscribe(pipe)
	for i := 0;i<10;i++ {
		pipe <- i
	}
}
