package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{} 
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		wg.Add(len(channels))
		result := make(chan interface{})
		for _, doneChan := range channels {
			go func ()  {
				defer wg.Done()
				<-doneChan
				result<-struct{}{}
			}()
		}
		go func() {
			defer close(result)
			wg.Wait()
		}()
		return result
	} 
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	// ch := make(chan interface{})
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v", time.Since(start))
}
