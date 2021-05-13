package main

import "fmt"

func main() {
	c := make(chan int)
	a := make(chan string)

	go foo(c, a)
	bar(c, a)

	fmt.Println("Program Biter")
}

func foo(x chan<- int, y chan<- string) {
	x <- 21
	y <- "Harun"
}
func bar(x <-chan int, y <-chan string) {
	fmt.Println(<-x, <-y)
}
