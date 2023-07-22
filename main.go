package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi!")
	horizontalDisplayLine := " ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----"

	current_number := 1
	fmt.Println(current_number)
	verticalDisplaySeparator := ""

	for i := 0; i <= 14; i++ {
		fmt.Println(horizontalDisplayLine)
		fmt.Println(verticalDisplaySeparator)
	}
	fmt.Println(horizontalDisplayLine)

	result1 := AddLeadingZerosToMakeIntegerThreeDigitsLong(0)
	result2 := AddLeadingZerosToMakeIntegerThreeDigitsLong(13)
	result3 := AddLeadingZerosToMakeIntegerThreeDigitsLong(890)
	result4 := AddLeadingZerosToMakeIntegerThreeDigitsLong(9090)
	fmt.Println(result1, result2, result3, result4)

}
