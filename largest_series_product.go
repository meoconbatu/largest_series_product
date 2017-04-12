package lsproduct

import (
	"errors"
	"strconv"
)

const testVersion = 4

func LargestSeriesProduct(digits string, span int) (product int64, err error) {
	product = -1
	digitsLen := len(digits)

	if span < 0 || span > digitsLen {
		err = errors.New("error")
		return
	}
	for i := 0; i+span <= digitsLen; i++ {
		seriesAtIndex := digits[i : i+span]
		tempProduct, tempErr := Multiple(seriesAtIndex)
		if tempErr != nil {
			err = tempErr
			return
		}
		if tempProduct > product {
			product = tempProduct
		}
	}
	return
}
func Multiple(series string) (output int64, err error) {
	output = 1
	var a int64
	for i := 0; i < len(series); i++ {
		if a, err = strconv.ParseInt(string(series[i]), 10, 64); err == nil {
			output = output * a
		} else {
			return
		}
	}
	return
}
