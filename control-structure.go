package main

import "fmt"

func main() {
	/*

		if condition {

		}else if another_condition {

		}else {

		}
	*/

	x := 71
	if x < 10 {
		fmt.Println("x is less than 10\n")
	} else {
		fmt.Println("X is greater than 10\n")
	}

	// go doesn't require parentheses around conditions
	// you can define variables inside the if statement

	// 2. Switch Case
	// Switch statement allows multi-way branching based on a value or expression. Unlike other languages,
	// Go's switch doesn't require break to terminate, and it supports matching expression directly.
	day := "saturday"
	switch day {
	case "monday":
		fmt.Println("Start of the week")
		// Block executed if expression matches value 1
	case "saturday":
		fmt.Println("End of the week")
	// Block executed if expression matched value 2
	default:
		fmt.Println("Mid-week going on")
		// Default executed if none of the cases match

	}

	// fallthrough:
	// GO allows the fallthrough keyword  to explicitly force execution to continue to the next case.
	switch day := "Monday"; day {
	case "Monday":
		fmt.Println("Start of the week")
		fallthrough
	case "Tuesday":
		fmt.Println("Second day of the week")
	}

	// Switch without an  expression:
	// you can use switch without a condition, making it work like a series of if/else statements.
	xx := 7
	switch {
	case xx > 10:
		fmt.Println("x is greater than 10")
	case xx == 10:
		fmt.Println("x is exactly 10")
	default:
		fmt.Println("x is less than 10")
	}

	// for loop
	// In go the for loop is the only loop construct. It can be used in three main ways: tradition loops, while loops, and infinite loops.

	// basic for loop:
	// this is tradition for loop(similar to other languages like c):
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop (Go for loop can act as like a while loop)
	fmt.Println("like while loop:")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	//	fmt.Println("This will run forever")
	//}

	// Range-based for loops:

	// go allows looping over arrays, slices, maps, and channels using range
	// Array/Slice Example:
	fmt.Println("Loop over the array/slice ")
	arr := []int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v) // index, value
	}

	// Map
	fmt.Println("Looping over the Map: ")
	myMap := map[int]string{71: "abhii", 69: "who?"}
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// Breaking and continuing loops:
	// you can control the flow of loops using break and continue:
	// break: Exists the loop completely
	// continue: Skip the  current iteration and move to next iteration
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // exit the loop
		}
		if i == 2 {
			continue // skip current iteration
		}
		fmt.Println(i)
	}
}
