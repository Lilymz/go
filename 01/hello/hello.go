package main

import (
	"fmt"
	greeting "greetings"
)

func main() {
	var hello string
	defer func() {
		if hello != "" {
			fmt.Println(hello)
		}
	}()
	greeting.Hello(hello)
}
