package main

import "runtime"

func main() {
	n := 7
	runtime.GOMAXPROCS(n) // // set number of OS threads that can execute Go code in parallel
}

// Concurrency vs Parallelism in Go
/*
Concept	        Description
Concurrency	Structuring your program to deal with multiple tasks at once, not necessarily simultaneously.
                Tasks may overlap in execution but don't truly run at the same time.
Parallelism	Executing multiple tasks literally at the same time (on multiple cores/CPUs). Full hardware-level execution in parallel.

// In Simple way
Concept	        Analogy
Concurrency	A single chef switching between dishes quickly.
Parallelism	Multiple chefs cooking dishes simultaneously on separate stoves.


*/

//  Does Go Support Both?
/*
Yes:

Go supports concurrency through goroutines and channels.
Go supports parallelism when GOMAXPROCS > 1 and multiple CPU cores are available.

Key Points:
Concurrency is about structure; parallelism is about execution.
You can have concurrency without parallelism (e.g., single-core CPU).
You can have parallelism without concurrency (e.g., SIMD instructions).
*/
