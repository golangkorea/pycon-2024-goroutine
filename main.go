package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello("Golang")
	sayHello("Korea")
}

func sayHello(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello, %s!\n", name)
		time.Sleep(500 * time.Millisecond)
	}
}
