package y946

import (
	"fmt"
	"testing"
)

func Test_solv(t *testing.T) {
	pushed := []int{2, 1, 0}
	//popped := []int {4,3,5,1,2}
	popped := []int{1, 2, 0}
	fmt.Println(validateStackSequences(pushed, popped))
}
