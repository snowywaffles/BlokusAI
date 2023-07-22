package main

import (
	"strconv"
)

func AddLeadingZerosToMakeIntegerThreeDigitsLong(integer int) string {
	intString := strconv.Itoa(integer)
	if len(intString) >= 4 {
		return "the integer input is four digits or longer"
	} else if len(intString) == 3 {
		// do nothing
	} else if len(intString) == 2 {
		intString = "0" + intString
	} else if len(intString) == 1 {
		intString = "00" + intString
	}
	return intString
}
