package main

import (
	"fmt"

	"github.com/victoryRo/learn-go-with-test/hello"
)

func main() {
	fmt.Println(hello.Hello("Victor", ""))
	fmt.Println(hello.Hello("", ""))
	fmt.Println(hello.Hello("Maria", "Spanish"))
	fmt.Println(hello.Hello("angela", "French"))
}
