package main

import (
	"fmt"
	"runtime"
	"sync"
)

var val int
var wg sync.WaitGroup

func main() {
	val = 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := val
			runtime.Gosched()
			v++
			val = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(val)
}
