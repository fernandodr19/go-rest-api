package main

import(
	"fmt"
	"time"
)

func main () {
	start := time.Now()

    channelX := make(chan string)
    channelY := make(chan string)

	go compute("x", 5, channelX)
	go compute("y", 5, channelY)

	for i := 0; i < 2; i++ {
        select {
        case msg1 := <-channelX:
            fmt.Println("received", msg1)
        case msg2 := <-channelY:
			fmt.Println("received", msg2)
		// case <-time.After(3 * time.Second):
		// 	fmt.Println("shutdown threads")
		// default:
		// 	fmt.Println("non blocking select")
        }
	}
	
	// I could just use syc.WaitGroup to wait a go routine to be finished
	// Also check Mutex


	elapsedTime := time.Since(start)
    fmt.Println("Execution took", elapsedTime)
}

func compute(msg string, value int, c chan<- string) {
	for i := 0; i <= value; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
	c <- msg+" done"
}