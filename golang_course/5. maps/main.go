package main

import "fmt"

func main() {
	// declaration
	// 1
	m := map[string]int{
		"blue":  1,
		"green": 2,
	}

	// 2
	n := make(map[int]int)
	n[10] = 1
	fmt.Println(n)

	//update
	m["red"] = 1
	fmt.Println(m)

	// iteration
	for color, number := range m {
		fmt.Println(color, number)
	}

	// delete
	delete(m, "green")
	fmt.Println(m)
}
