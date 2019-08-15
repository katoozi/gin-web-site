package templatefuncs

import (
	"fmt"
	"strconv"

	"github.com/dustin/go-humanize"
)

// IntComma take given number and return string on number with thousand seperator
// example => CommaSeperated(12000) -> 12,000
func IntComma(number int) string {
	return fmt.Sprintf("%s", humanize.Comma(int64(number)))
}

// NumberToWord will convert 0 to 9 number to english represent
func NumberToWord(number int) string {
	if number >= 10 {
		return strconv.Itoa(number)
	}
	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	return words[number-1]
}
