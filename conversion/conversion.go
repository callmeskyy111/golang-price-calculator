package conversion

import (
	"errors"
	"strconv"
)

// Outsourcing shareable conversion-logic

func StringsToFloats(strs []string) ([]float64, error) {
	var floats []float64
	for _, str := range strs {
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float! :(")
		}
		floats = append(floats, floatVal)
	}
	return  floats,nil
}