package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var (
	foo           = false
	maxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (aPrime, bPrime string) {
	aPrime = b
	bPrime = a
	return
}

func printVar(a interface{}) {
	fmt.Printf("Type: %T Value: %v\n", a, a)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("hello, world")

	fmt.Println("The time is", time.Now(), "hi")

	fmt.Println("My favorite number is", rand.Intn(10))

	val := math.Pow(float64(rand.Intn(10)), 2)
	// fval := float64(val)
	fmt.Println("The square root of", val, "is", math.Sqrt(val))

	fmt.Println("The best number is", math.Pi)

	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Println("The sum of", x, "and", y, "is", add(x, y))

	a := "Cat"
	b := "Dog"
	aPrime, bPrime := swap(a, b)
	fmt.Println("If you swap", a, "with", b, "\b, you get", aPrime, "and", bPrime)

	printVar(foo)
	printVar(maxInt)
	printVar(z)

	var someVar float64 = 3.1415
	var someVarInt int = int(someVar)
	fmt.Println("Floating point value", someVar, "and integer value", int(someVarInt))
}
