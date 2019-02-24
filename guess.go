package main

import (
	"bufio" // function that has something to do with recording input from console
	"fmt"   // imput output function
	"os"    // a library that allows your application to make system calls
	"time"
)

// everything is inside the scope of this main function
func main() {
	// assigning a variable labelled scanner. Easier to type than bufio.NewScanner blabla
	scanner := bufio.NewScanner(os.Stdin)

	//Game introduction
	fmt.Println("Please think of a number between 1 and 100...")
	fmt.Println("Press ENTER when you are ready!")
	// recording input from the player - basically, press Enter.
	scanner.Scan()

	// the variable low is currently '1' and the variable high is currently '100'
	low := 1
	high := 100

	// below is a loop which will complete a set of instructions
	for {
		// below, guess is currently the variable of (low + high)/2 so right now it's going to be 101/2 = 50. (NOT REAL MATHS it should be 50.5 right?)
		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess) // I guess the number is 50
		fmt.Println("Is that:")
		fmt.Println("(a): too high?")
		fmt.Println("(b): too low?")
		fmt.Println("or")
		fmt.Println("(c): correct?")
		// Now waiting for user to input some stuff, hopefully a, b or c.
		scanner.Scan()
		// response is currently the variable for whatever string is typed into the console i.e a, b or c.
		response := scanner.Text()

		// If the response is a, then high will become 49, if b then low will be 2. If c then print then exit loop.
		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I win.")
			break
		} else {
			fmt.Println("Choose a, b or c please.")
		}
	}
	duration := time.Second // variable 'duration' is currently one second
	time.Sleep(duration)    // sleep for 1 second
	fmt.Println("Goodbye!")
	time.Sleep(duration) // sleep for 1 second
}
