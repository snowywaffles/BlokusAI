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
	// define colors
	cyan := color.New(color.FgCyan)
	magenta := color.New(color.FgMagenta)

	fmt.Println()
	fmt.Println(horizontalDisplayLine) // to make the grid look nice (for formatting purposes)

	for i := 0; i <= 13; i++ {
		fmt.Printf("|")
		for j := 0; j <= 13; j++ {
			if valueInCell[cellNumber] == "B" { // support for custom colors!
				cyan.Printf(" %-4s", valueInCell[cellNumber])
			} else if valueInCell[cellNumber] == "P" {
				magenta.Printf(" %-4s", valueInCell[cellNumber])
			} else {
				fmt.Printf(" %-4s", valueInCell[cellNumber])
			}
			fmt.Printf(" |")
			cellNumber++
		}
		fmt.Println()
		fmt.Println(horizontalDisplayLine)
	}
	return nil
}

func main() {

	valueInCell := make(map[int]string)
	valueInCell[1] = "B"
	valueInCell[14] = "P"
	valueInCell[183] = "B"
	valueInCell[196] = "P"

	err := draw14x14GridWithValues(valueInCell)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(valueInCell)
	color.Cyan(valueInCell[1])
}
