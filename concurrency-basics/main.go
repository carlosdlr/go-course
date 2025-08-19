package main

import (
	"fmt"
	"time"
)

func greet(phrase string, donechan chan bool) {
	fmt.Println("Hello!", phrase)
	donechan <- true // signal that the goroutine is done using a channel and <-
}

func slowGreet(phrase string, donechan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	donechan <- true // signal that the goroutine is done using a channel and <-
	close(donechan)  // close the channel to signal that no more values will be sent, this can be used when we know
	// that no more goroutines will send values to this channel or it's the longest running goroutine
	// this is useful to avoid deadlocks in the main function when waiting for all goroutines
	// to finish
	// if we don't close the channel, the main function will block forever waiting for a value
	// from the channel, because it will never receive a value from the channel again
}

func main() {
	//dones := make([]chan bool, 4) // buffered channel with a capacity of
	done := make(chan bool) // create a channel to keep the main function running and to check when goroutines finish

	//dones[0] = make(chan bool)
	go greet("Nice to meet you!", done) // go routine works in the background in a non-blocking way
	//dones[1] = make(chan bool)
	go greet("How are you?", done)
	//dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	//dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)
	/*<-done // block here until we get a value from the done channel
	<-done
	<-done
	<-done // wait for all 4 goroutines to finish
	*/
	// for _, done := range dones {
	// 	<-done // wait for all 4 goroutines to finish
	// }
	for range done { // wait for all goroutines to finish
		//fmt.Println("A goroutine has finished its work.")
	}
}
