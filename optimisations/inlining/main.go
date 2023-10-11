package main

import "fmt"

func main() {
	a1()
}

func a1() {
	a2()
}

func a2() {
	for i := 0; i < 10; i += 2 {
		if i%3 == 0 {
			fmt.Println("hello world!")
		}
	}
}
