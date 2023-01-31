package main

import (
	"fmt"
	"strings"
)

func factorial(n int, ch chan int) {
	var f = 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	ch <- f
}

func main() {
	ch := make(chan int)

	defer close(ch)

	//C1: FUNCTION
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}

	fmt.Println(strings.Repeat("-", 30))

	//C2: ANONYMOUS FUNCTION
	for i := 1; i <= 20; i++ {
		go func(n int, c chan int) {
			var f = 1
			for i := 2; i <= n; i++ {
				f *= i
			}

			c <- f
		}(i, ch)

		fmt.Printf("The factorial of %d is %d\n", i, <-ch)
	}
}
