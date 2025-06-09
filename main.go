package main

import "fmt"

func main() {
	BB()
}

func BB() {
	fmt.Print(1)
}

func BB() string {
	return "hello"
}

func CC() {
	fmt.Println("hello")
}
