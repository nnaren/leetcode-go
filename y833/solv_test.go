package y833

import (
    "testing"
)

func Test_solv(t *testing.T) {
    s := "abcd"
    indices := []int{0, 2}
    sources := []string{"ab", "ec"}
    targets := []string{"eee", "fff"}
    findReplaceString(s, indices,sources,targets)

}
