package main

import (
	"fmt"
	"strings"
)

func main() {
	splitMe := "[some string] to extract"
	split := strings.SplitN(splitMe, "]", 2)
	fmt.Printf("Split string: '%s'\n", strings.TrimSpace(split[1]))

	fmt.Println("Hello, world...")
	variables()
}

func variables() {
	x := 10
	fmt.Printf("x: %d\n", x)

	x = 15
	fmt.Printf("x: %d\n", x)

	// Not allowed
	//x := 20
	//fmt.Printf("x: %d\n", x)

	// allowed `:=`
	// as long as there is one new variable on the lefthand side of := then any of the other variables can already exist
	x, y := 20, 25
	fmt.Printf("x: %d, y: %d\n", x, y)
}
