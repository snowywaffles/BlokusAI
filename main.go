package main

import (
	"fmt"
	// "github.com/fatih/color"
	// "log"
)

func draw14x14GridWithValues(valueInCell map[int]string) error {
	cellNumber := 1
	horizontalDisplayLine := "---------------------------------------------------------------------------------------------------"

	fmt.Println()
	fmt.Println(horizontalDisplayLine) // to make the grid look nice (formatting purposes)

	for i := 0; i <= 13; i++ {
		fmt.Printf("|")
		for i := 0; i <= 13; i++ {
			fmt.Printf(" %-4s |", valueInCell[cellNumber])
			cellNumber++
		}
		fmt.Println()
		fmt.Println(horizontalDisplayLine)
	}
	return nil
}

func main() {

	valueInCell := make(map[int]string)
	valueInCell[1] = "W"
	valueInCell[10] = "bo"
	valueInCell[100] = "1111"

	draw14x14GridWithValues(valueInCell)
}
