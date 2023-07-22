package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi!")
	horizontalDisplayLine := " ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----"
	verticalDisplaySeparator := "|    |    |    |    |    |    |    |    |    |    |    |    |    |    |    |   "
	for i := 0; i <= 14; i++ {
		fmt.Println(horizontalDisplayLine)
		fmt.Println(verticalDisplaySeparator)
	}
	fmt.Println(horizontalDisplayLine)
}
