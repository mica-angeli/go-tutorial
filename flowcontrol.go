// Test code from Go tutorial at https://tour.golang.org/flowcontrol/
package main

import "fmt"

func clamp(x, min, max int) int {
	if x < min {
		return min
	} else if x > max {
		return max
	} else {
		return x
	}
}

func doublerLimited(x, lim int) int {
	if x = 2 * x; x < lim {
		return x
	}
	return lim
}

func doubler(x int) int {
	// This is Go's "while" loop
	for x < 10000 {
		x = 2 * x
	}
	return x
}

func sum() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println("The sum is", sum())

	fmt.Println("The number 3 doubled continuously is", doubler(3))

	fmt.Println("The number 4 doubled with a limit of 10 is", doublerLimited(4, 10))
	fmt.Println("The number 7 doubled with a limit of 10 is", doublerLimited(7, 10))

	// Clamp
	fmt.Println("The number 2 clamped is", clamp(2, 0, 10))
	fmt.Println("The number -2 clamped is", clamp(-2, 0, 10))
	fmt.Println("The number 2000 clamped is", clamp(2000, 0, 10))
}
