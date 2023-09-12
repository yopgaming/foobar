package foobar

import "fmt"

func Foo(name string) string {

	message := fmt.Sprintf("Foobar ! %v", name)
	return message
}
