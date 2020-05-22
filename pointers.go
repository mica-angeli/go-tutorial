package main

import "fmt"

func foo() *[]int {
	i := []int{1, 2, 3}
	return &i
}

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	p_foo := foo()
	fmt.Println((*p_foo)[1])
}
