package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "Harun",
		last:  "Unal",
		age:   21,
	}
	saySomething(&p1) // pointer olması lazım DİKKAT
}

func (r *person) speak() {
	fmt.Println("Hello", r.first)
}

func saySomething(h human) {
	h.speak()
}
