package conversion

import (
	"errors"
	"strconv"
	"strings"
)

func StringsToFloats(strs []string) ([]float64, error) {
	var floats []float64
	for _, str := range strs {
		str = strings.TrimSpace(str) // Removes spaces, tabs, newlines
		if str == "" {
			continue // Skip empty lines
		}
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float: " + err.Error())
		}
		floats = append(floats, floatVal)
	}
	return floats, nil
}
