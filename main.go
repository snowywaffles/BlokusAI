package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func main() {

	valueInCell := make(map[int]string)
	valueInCell[1] = "W"
	valueInCell[10] = "bo"
	valueInCell[100] = "1111"

	cellNumber := 1
	horizontalDisplayLine := "---------------------------------------------------------------------------------------------------"

	fmt.Println()
	fmt.Println(horizontalDisplayLine) // to make the grid look nice

	for i := 0; i <= 13; i++ {
		fmt.Printf("|")
		for i := 0; i <= 13; i++ {
			if len(valueInCell[cellNumber]) > 4 {
				log.Panic("the value in (cell: ", cellNumber, ") is longer than 4 characters")
			}
			fmt.Printf(" %-4s |", valueInCell[cellNumber])
			cellNumber++
		}
		fmt.Println()
		fmt.Println(horizontalDisplayLine)
	}

	color.Cyan(horizontalDisplayLine)
	fmt.Println(valueInCell)
}
