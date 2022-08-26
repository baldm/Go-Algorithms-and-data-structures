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

func QuickGraphInputParser() (int, [][]int) {
	// A faster version of the graph parser using buffer
	var nodeAmount int
	var edgeAmount int
	var firstCord int
	var secondCord int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d\n", &nodeAmount)
	fmt.Fscanf(reader, "%d\n", &edgeAmount)

	// Build edge matrix
	edgesMatrix := make([][]int, edgeAmount)
	// Parse input into edge matrix
	for i := 0; i < edgeAmount; i++ {
		fmt.Fscanf(reader, "%d %d\n", &firstCord, &secondCord)
		edgesMatrix[i] = []int{firstCord, secondCord}
	}

	return nodeAmount, edgesMatrix
}

func GraphQuestionParser() ([][]int, []int) {
	var Q1Amount int
	var Q2Amount int
	var firstCord int
	var secondCord int

	reader := bufio.NewReader(os.Stdin)

	// Build neighbor question matrix
	fmt.Fscanf(reader, "%d\n", &Q1Amount)
	Q1List := make([][]int, Q1Amount)
	// Parse input into edge matrix
	for i := 0; i < Q1Amount; i++ {
		fmt.Fscanf(reader, "%d %d\n", &firstCord, &secondCord)
		Q1List[i] = []int{firstCord, secondCord}
	}

	// Build neighbor question matrix
	fmt.Fscanf(reader, "%d\n", &Q2Amount)
	Q2List := make([]int, Q2Amount)
	// Parse input into edge matrix
	for i := 0; i < Q2Amount; i++ {
		fmt.Fscanf(reader, "%d\n", &firstCord)
		Q2List[i] = firstCord
	}
	return Q1List, Q2List
}
