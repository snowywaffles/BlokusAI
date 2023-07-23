package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func draw14x14GridWithValues(valueInCell map[int]string) error {

	// A standard Blokus Duo board is 14x14, or 196 squares.
	// cellNumber begins at 1, and is incremented every time the next square is made.
	// TOP LEFT: cellNumber = 1
	// TOP RIGHT: cellNumber = 14
	// BOTTOM LEFT: cellNumber = 183
	// BOTTOM RIGHT: cellNumber = 196

	cellNumber := 1
	horizontalDisplayLine := "---------------------------------------------------------------------------------------------------"

	fmt.Println()
	fmt.Println(horizontalDisplayLine) // to make the grid look nice (for formatting purposes)

	cyan := color.New(color.Cyan)
	cyan.Printf(" %-4s |", valueInCell[cellNumber])

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
	valueInCell[14] = "bo"
	valueInCell[183] = "hey"
	valueInCell[196] = "1111"

	err := draw14x14GridWithValues(valueInCell)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(valueInCell)
	color.Cyan(valueInCell[1])
}
