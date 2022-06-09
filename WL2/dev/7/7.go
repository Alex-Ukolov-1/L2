package main

import (
	"fmt"
	"time"
)

// or для каждого канала запускается горутина, которая закрывает основной канал если дочерний закрылся
func or(channels ...<-chan interface{}) <-chan interface{} {
	singleChan := make(chan interface{})
	for _, value := range channels {
		go func(ch <-chan interface{}) {
			select {
			case _, opened := <-ch:
				if !opened {
					close(singleChan)
				}
			}
		}(value)
	}
	return singleChan
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("foneafter %v", time.Since(start))
}
