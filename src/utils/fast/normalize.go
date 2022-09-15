package Fast

import (
	"fmt"
	"unicode"
)

func appendData(rangeStr string) string {

	var str bool
	var num bool
	var convert string

	if rangeStr != "" {

		for _, char := range rangeStr {
			if unicode.IsNumber(char) {
				num = true
			} else {
				str = true
			}
		}

		if num && str {
			convert = "[si]"
			return convert
		}

		if num {
			convert = "[i]"
		}

		if str {
			convert = "[s]"
		}

	}

	return convert
}

func normalize(url string) (string, error) {

	var lenUrl = len(url)
	var rangeStr string
	var delimeter = 0
	var str string

	for pos, char := range url {

		if char == '/' {

			delimeter++
			str += appendData(rangeStr)

			rangeStr = ""
		} else {
			rangeStr += string(char)
		}

		if pos == lenUrl-1 && char != '/' {

			str += appendData(rangeStr)
		}
	}

	if delimeter == 1 && rangeStr == "" {
		return "/", nil
	}

	fmt.Println(str)

	return "", nil
}
