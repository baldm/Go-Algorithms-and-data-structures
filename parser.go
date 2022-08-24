//This module parsers the different input each exercise needs.
// CodeJudge inputs strings into a virtual terminal like a regular terminal

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TwoLines() []int {

	var length int

	// Takes length of array
	fmt.Scanln(&length)
	// Create array
	var inputSlice = make([]int, length)

	// Create reader and read string with spaces
	in := bufio.NewReader(os.Stdin)
	inputString, _ := in.ReadString('\n')
	// Convert to slice
	numbers := strings.Fields(inputString)

	for i := 0; i < length; i++ {
		inputSlice[i], _ = strconv.Atoi(numbers[i])
	}

	return inputSlice
}

func FourLines() ([]int, []int) {
	return TwoLines(), TwoLines()
}
