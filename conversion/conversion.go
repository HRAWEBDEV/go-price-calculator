package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {

	floats := make([]float64, len(strings))
	for i, line := range strings {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("converting price to float failed")
			fmt.Println(err)
			return nil, errors.New("could not convert price to float")
		}

		floats[i] = floatPrice
	}
	return floats, nil
}
