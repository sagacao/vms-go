package utils

import "strconv"

func ToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return value
}

func ToFloat64(str string) float64 {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return float64(value)
}

func ToString(value int) string {
	return strconv.FormatUint(uint64(value), 10)
}
