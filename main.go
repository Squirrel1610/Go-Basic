package main

import (
	"fmt"
	"math"
)

func main() {
	ch := make(chan float64)

	for i := 1.; i <= 50.; i++ {
		go func(n float64, c chan float64) {
			c <- math.Pow(n, 2)
		}(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-ch)
	}
}
