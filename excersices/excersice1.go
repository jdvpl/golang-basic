package excersices

import "strconv"

func ConvertToNumber(text string) (int, string) {
	num, err := strconv.Atoi(text)

	if err != nil {
		return 0, "Not a number " + err.Error()
	}
	if num > 1000 {
		return num, "Greater than 1000"
	} else {
		return num, "Less than 1000"
	}
}
