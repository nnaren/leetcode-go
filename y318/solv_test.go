package y318

import (
	"fmt"
	"testing"
)

func Test_solv(t *testing.T) {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println(maxProduct(words))
}
