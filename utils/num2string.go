package utils

import (
	"sort"
	"strconv"
)

func num2Bytes(num int) []byte {
	return []byte(strconv.Itoa(num))
}

func num2String(num int) string {
	return strconv.Itoa(num)
}

func numBytesSort(num []byte) []byte {
	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	return num
}
