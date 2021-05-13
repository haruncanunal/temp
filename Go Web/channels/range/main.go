package main

import "fmt"

func main() {
	c := make(chan int)
	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

}
func foo(x chan<- int) {
	for i := 0; i < 100; i++ {
		x <- i
	}
	close(x)
}
