package main

import (
	"log"
	"net/http"
	"os"

	"github.com/victoryRo/learn-go-with-test/dependency"
)

func main() {
	// fmt.Println(hello.Hello("Victor", ""))
	// fmt.Println(hello.Hello("", ""))
	// fmt.Println(hello.Hello("Maria", "Spanish"))
	// fmt.Println(hello.Hello("angela", "French"))

	dependency.Greet(os.Stdout, "Love")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dependency.Greet(w, "world")
}
