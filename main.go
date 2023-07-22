package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	current_number := 1
	horizontalDisplayLine := ""
	for i := 0; i <= 13; i++ {
		verticalDisplaySeparator := "|"
		horizontalDisplayLine = ""
		for i := 0; i <= 13; i++ {
			verticalDisplaySeparator = verticalDisplaySeparator + " " + (AddLeadingZerosToMakeIntegerThreeDigitsLong(current_number)) + " |"
			horizontalDisplayLine = horizontalDisplayLine + " -----"
			current_number++
		}
		fmt.Println(horizontalDisplayLine)
		fmt.Println(verticalDisplaySeparator)
	}
	color.Cyan(horizontalDisplayLine)
}
