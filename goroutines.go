package main

import (
	"fmt"
	"time"
)

func wait() {
	fmt.Println("wait started")
	time.Sleep(5 * time.Second)
	fmt.Println("wait finished")
}

func main() {
	defer fmt.Println("done")
	go wait()
	time.Sleep(6 * time.Second)
}
