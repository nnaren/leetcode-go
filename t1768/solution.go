git package t1768

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	n := len(students)
	oneStuNum := 0
	for _, k := range students {
		if k == 1 {
			oneStuNum++
		}
	}
	oneSdwchNum := 0
	for _, k := range sandwiches {
		if k == 1 {
			oneSdwchNum++
		}
	}
	if oneStuNum == oneSdwchNum {
		return 0
	} else if oneStuNum > oneSdwchNum {
		zeroNum := n - oneStuNum
		for idx, i := range sandwiches {
			if i == 0 {
				zeroNum--
				if zeroNum < 0 {
					return n - idx
				}

			}
		}
	} else {
		for idx, i := range sandwiches {
			if i == 1 {
				oneStuNum--
				if oneStuNum < 0 {
					return n - idx
				}
			}
		}
	}
	return 0
}

func main() {
	//student := []int {1, 1, 1, 0, 0, 1}
	//sandwiches := []int{1,0,0, 0,1,1}
	student := []int{1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 0, 0, 0}
	sandwiches := []int{0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 0, 1, 0}
	ans := countStudents(student, sandwiches)
	fmt.Println("无法吃午餐的人数:", ans)
}
