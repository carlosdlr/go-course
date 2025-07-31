package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64
	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Fail to convert string to float at index " + strconv.Itoa(stringIndex) + ": " + err.Error())
		}
		floats = append(floats, floatPrice)
	}

	return floats, nil
}
