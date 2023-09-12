package foobar

import "fmt"

func Bar(name string) string {

	message := fmt.Sprintf("Go to a bar and meet people, %v", name)
	return message
}
