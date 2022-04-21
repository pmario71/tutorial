package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Channel_for_sync(t *testing.T) {
	ch := make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second * 2)
}

func send(ch chan int) {
	ch <- 1
	fmt.Println("Sending value to channel complete")
}

func receive(ch chan int) {
	time.Sleep(time.Second * 1)
	fmt.Println("Timeout finished")
	received := <-ch
	fmt.Println("Received: ", received)
}

func Test_WaitGroup(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go SyncOnWaitGroup(wg)

	wg.Wait()
}

func SyncOnWaitGroup(wg *sync.WaitGroup) {
	wg.Done()
}
