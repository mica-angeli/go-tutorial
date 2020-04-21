package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func sqrt(x float64, precision int) float64 {
	// Special cases
	if math.IsNaN(x) {
		return math.NaN()
	} else if math.IsInf(x, 1) {
		return math.Inf(1)
	} else if x < 0.0 {
		return math.NaN()
	} else if x == 0.0 {
		return 0.0
	}

	z := 1.0
	for i := 0; i < precision; i++ {
		diff := x - z*z
		if math.Abs(diff) == 0.0 {
			break
		}
		z += diff / (2 * z)
	}
	return z
}

func boolToResult(b bool) (result string) {
	if b {
		result = "SUCCESS"
	} else {
		result = "FAILURE"
	}
	return
}

func testSqrt(useBigNumbers bool, precision int) bool {
	var n int
	if useBigNumbers {
		n = rand.Int()
	} else {
		n = rand.Intn(100)
	}
	x := float64(n * n)
	fmt.Print("\tTesting square root of ", x, " :\t")
	stdResult := math.Sqrt(x)
	myResult := sqrt(x, precision)

	var testResult bool
	if math.IsNaN(stdResult) {
		testResult = math.IsNaN(myResult)
	} else {
		testResult = stdResult == myResult
	}

	fmt.Println(boolToResult(testResult))
	return testResult
}

func testPrecision(precision int) bool {
	for i := 0; i < 1000000; i++ {
		result := testSqrt(true, precision)
		if !result {
			return false
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var mode string
	flag.StringVar(&mode, "mode", mode, "Name of test to run.  Choices are small, big, precision")
	flag.Parse()

	switch mode {
	case "small":
		fmt.Println("Testing 10 small numbers:")
		for i := 0; i < 10; i++ {
			testSqrt(false, 37)
		}
	case "big":
		fmt.Println("Testing 10 big numbers:")
		for i := 0; i < 10; i++ {
			testSqrt(true, 37)
		}
	case "precision":
		fmt.Println("Finding minimum level of precision:")
		precision := 36
		for ; !testPrecision(precision); precision++ {
		}
		fmt.Println("Minimum level of precision necessary", precision)
	default:
		panic("Invalid mode argument")
	}

}
