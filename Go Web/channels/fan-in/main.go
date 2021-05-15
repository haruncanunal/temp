package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("errors.txt")
	if err != nil {
		fmt.Println("err happened", err)
	}
	defer f.Close()

	log.SetOutput(f) // bununla log çıkışını belirledik

	f2, err := os.Open("hcu.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f2.Close()
	fmt.Println("Dosyaya bak")
}
