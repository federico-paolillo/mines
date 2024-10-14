package validation

import "strconv"

var uppercaseY = "Y"
var lowercaseN = "n"

func IsYN(value string) bool {
	return IsY(value) || IsN(value)
}

func IsY(value string) bool {
	return value == uppercaseY
}

func IsN(value string) bool {
	return value == lowercaseN
}

func IsNumber(value string) bool {
	_, err := strconv.Atoi(value)

	return err == nil
}

func ToNumberUnsafely(value string) int {
	i, _ := strconv.Atoi(value)

	return i
}
