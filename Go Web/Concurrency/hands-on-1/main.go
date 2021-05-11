package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go boo()
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	fmt.Println("Benim adım harun")
	wg.Done()
}

func boo() {
	fmt.Println("Benim adım boo")
	wg.Done()
}
