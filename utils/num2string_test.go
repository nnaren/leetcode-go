package utils

import (
	"fmt"
	"testing"
)

func Test_solv(t *testing.T) {

	fmt.Println(num2Bytes(4325))

}
func Test_solv2(t *testing.T) {

	fmt.Println(num2String(4325))

}

func Test_solv3(t *testing.T) {

	fmt.Println(numBytesSort(num2Bytes(4325)))
	fmt.Println(string(numBytesSort(num2Bytes(4325))))

}
