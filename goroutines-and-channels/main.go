package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepy(done chan struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	time.Sleep(time.Second * 5)
}

func checkDone(done chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-done:
			fmt.Printf("we are done!\n")
			wg.Done()
			return
		}
	}
}

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go checkDone(done, &wg)
	go sleepy(done)
	wg.Wait()
	fmt.Printf("main function done\n")
}
