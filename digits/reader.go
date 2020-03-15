package digits

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// SerializeDigits retrieves the digits contained in a file and shrinks it to the number of
// digits wanted.
func SerializeDigits(filePath string, digits int) string {
	fmt.Printf("Reading %s...\n", filePath)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		err := fmt.Sprintf("Cannot read file %s.", filePath)
		panic(err)
	}

	content := strings.Replace(string(data), ".", "", -1)

	// Change digits number if they are bigger than the number of digits read from the file
	if digits > len(content) {
		fmt.Println("Too much digits are asked, returning the maximum number of digits by default.")
		digits = len(content)
	}

	return content[:digits-1]
}