package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	c := make(chan int)

	//receive
	c1 := make(<-chan int)

	//send
	c2 := make(chan<- int)

	fmt.Printf("%T, %T, %T\n", c, c1, c2)

	go f1(4, c)

	n := <-c

	fmt.Println(n)
}
