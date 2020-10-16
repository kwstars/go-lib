package main

import "fmt"

type Foo struct {
	S string
	N int
	F float64
}

func getS(foo Foo) string {
	// Bad! Use wire.FieldsOf instead.
	return foo.S
}

func provideFoo() Foo {
	return Foo{S: "Hello, World!", N: 1, F: 3.14}
}

func main() {
	i := injectedMessage()
	fmt.Println(i)
}
