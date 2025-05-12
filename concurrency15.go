package main

import (
	"fmt"
	"time"
)

// Concurrency
// Go makes concurrency simple with goroutines and channels.
// 1.  Goroutines: Lightweight Threads
// A goroutine is a function running concurrently with other functions.
// It is extremely lightweight (2KB stack initially) and managed by the Go runtime.
// go sayHello()
// This launches the function sayHello() as a new goroutine

// 2. Channels: Communication Between Goroutines
/*
	Channels let goroutines communicate safely by passing values.
	ch := make(chan int)
	ch <- 5       // send
	val := <-ch   // receive
*/

// goroutines
func sayHello() {
	fmt.Println("Hello from goroutine")
}

// channels:
func worker(ch chan string) {
	ch <- "Hello from worker" //send
}

/*
Channels are typed (e.g., chan int, chan string)
Sending to a channel blocks until another goroutine receives from it (unless buffered)
Buffered channels let you send without blocking:

ch := make(chan int, 2)
ch <- 1
ch <- 2

Edge Cases:
Deadlock: If both sender and receiver are waiting.
Race conditions: If shared variables are accessed without proper synchronization.
Closing channels: Use close(ch) to signal no more values will be sent.
*/

func main() {
	go sayHello()                      // concurrent
	fmt.Println("Main done")           // may finish before goroutine runs
	time.Sleep(500 * time.Microsecond) // give time for goroutine
	//  If main() finishes before sayHello(), the program may exit before printing anything.
	// That's why we use time.Sleep() or channels to synchronize.

	// channels:
	ch := make(chan string)
	go worker(ch)
	msg := <-ch // receive
	fmt.Println(msg)
}
