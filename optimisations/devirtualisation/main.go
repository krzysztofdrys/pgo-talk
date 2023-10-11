package main

import "fmt"

type foo interface {
	bar()
}

type Foo struct{}

func (f Foo) bar() {
	fmt.Println("Hello world!")
}

func main() {
	fooer(Foo{})
}

func fooer(f foo) {
	f.bar()
}
