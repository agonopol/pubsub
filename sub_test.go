package pubsub

import (
	"log"
	"testing"
	"time"
)

func YUNOTestTwoSubscribers(t *testing.T) {
	sub := Subscriber()
	sub.AddSubscriber(func(i interface{}) {
		log.Printf("First subscriber consumed %v", i)
		time.Sleep(time.Millisecond)
	})
	sub.AddSubscriber(func(i interface{}) {
		log.Printf("Second subscriber consumed %v", i)
		time.Sleep(time.Millisecond)
	})
	pipe := make(chan interface{})
	sub.Subscribe(pipe)
	for i := 0; i < 10; i++ {
		pipe <- i
	}
}

func TestTwoPubTwoSub(t *testing.T) {
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
	pipe := make(chan interface{},10)
	sub := Subscriber()
	sub.AddSubscriber(func(i interface{}) {
		time.Sleep(time.Millisecond)
		log.Printf("First subscriber consumed %v", i)
		time.Sleep(time.Millisecond)
	})
	sub.AddSubscriber(func(i interface{}) {
		log.Printf("Second subscriber consumed %v", i)
		time.Sleep(time.Millisecond)
	})
	sub.Subscribe(pipe)
	pub.Publish(pipe)
}
