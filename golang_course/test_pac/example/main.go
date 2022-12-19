package main

import (
	"fmt"
	greeter "greet"
)

func main() {
	greeter.Hello()
	fmt.Println(greeter.Shark)

	c := greeter.Candy{"oct", "red"}
	fmt.Println(c)

	c.Reset()
	fmt.Println(c)

}
