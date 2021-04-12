package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	r := 0

	for _, v := range a {
		r += v
	}

	res <- r
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	fmt.Println("The sum of the array a is ", <-ch + <-ch)
	return -1
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(ConcurrentSum(a))
}
