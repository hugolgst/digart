package digits

import (
	"fmt"
	"strconv"
)

// ParseDigits returns a 3D array of integers which represents the segments of the first circle
// then the different lines of a segment and finally the points which represents the decimals.
func ParseDigits(number string) [][][]int {
	fmt.Printf("Parsing the digits\n")

	// Generate the 3D array
	var digits = make([][][]int, 10)

	// Iterate all the digits of the given number
	for digitIndex := 0; digitIndex < len(number)-1; digitIndex++ {
		previous, _ := strconv.Atoi(string(number[digitIndex]))
		current, _ := strconv.Atoi(string(number[digitIndex+1]))
		segments := digits[previous]

		// Create a new segment if no one has been found
		if len(segments) == 0 {
			segments = append(segments, MakeIntArray(10))
		}

		// Set the position of the digit in the given number for an array of 10 numbers
		position := digitIndex % 10

		// Then generate the segment with the current value and its position
		digits[previous] = GenerateSegment(segments, position, current)
	}

	return digits
}

// GenerateSegment returns a segment where the value has been filled
func GenerateSegment(segments [][]int, position, value int) [][]int {
	filled := false
	// Iterate the segments to find an available position
	for i := range segments {
		// If the position is free then fill it with the current digit
		if segments[i][position] == -1 {
			segments[i][position] = value
			filled = true
		}
	}

	// If a position has not been found in the existent segments' lines then create one
	// and fill it with the current digit
	if !filled {
		segments = append(segments, MakeIntArray(10))
		segments[len(segments)-1][position] = value
	}

	return segments
}

// MakeIntArray returns a `size`-elements-array which elements are -1
func MakeIntArray(size int) (array []int) {
	array = make([]int, size)
	for i := range array {
		array[i] = -1
	}

	return
}