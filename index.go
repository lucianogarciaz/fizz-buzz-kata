package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 100; i += 1 {
		object := Factory(i)
		if object != nil {
			fmt.Println(object.GetName())
		} else {
			fmt.Println(i)
		}
	}
}
