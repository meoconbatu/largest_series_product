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
	digitsLen := len(digits)

	if span < 0 || span > digitsLen || from < 0 || to < 0 {
		productChan <- product
		errChan <- errors.New("error")
		return
	}
	for i := from; i+span <= to; i++ {
		seriesAtIndex := digits[i : i+span]
		tempProduct, tempErr := Multiple(seriesAtIndex)
		if tempErr != nil {
			errChan <- errors.New("error")
			productChan <- product
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
