package main

import (
	"fmt"
)

func main() {
	current_number := 1
	horizontalDisplayLine := " ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----"
	for i := 0; i <= 13; i++ {
		verticalDisplaySeparator := "|"
		for i := 0; i <= 13; i++ {
			verticalDisplaySeparator = verticalDisplaySeparator + " " + AddLeadingZerosToMakeIntegerThreeDigitsLong(current_number) + " |"
			current_number++
		}
		fmt.Println(horizontalDisplayLine)
		fmt.Println(verticalDisplaySeparator)
	}
	fmt.Println(horizontalDisplayLine)
}
