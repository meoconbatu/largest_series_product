package lsproduct

import (
	"errors"
)

const testVersion = 4

func LargestSeriesProduct(digits string, span int) (product int64, err error) {
	product = -1
	digitsLen := len(digits)

	if span < 0 || span > digitsLen {
		err = errors.New("error")
		return
	}
	productChan := make(chan int64, digitsLen)
	errChan := make(chan error, 2)
	temp := int(digitsLen / 2)
	go Worker(productChan, errChan, digits, span, 0, temp)
	go Worker(productChan, errChan, digits, span, temp+1-span, digitsLen)
	for i := 0; i < 2; i++ {
		select {
		case err = <-errChan:
			return
		case tempProduct := <-productChan:
			if tempProduct > product {
				product = tempProduct
			}
		}
	}
	return
}
func Worker(productChan chan int64, errChan chan error, digits string, span int, from int, to int) {
	product := int64(-1)
	for i := from; i+span <= to; i++ {
		seriesAtIndex := digits[i : i+span]
		tempProduct, tempErr := Multiple(seriesAtIndex)
		if tempErr != nil {
			errChan <- errors.New("error")
			return
		}
		if tempProduct > product {
			product = tempProduct
		}
	}
	productChan <- product
}
func Multiple(series string) (output int64, err error) {
	output = 1
	for i := 0; i < len(series); i++ {
		c := series[i]
		if c < '0' || c > '9' {
			return 0, errors.New("not a number")
		}
		output *= int64(c - 48)
	}
	return
}
