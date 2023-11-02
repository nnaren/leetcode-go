package y2578

func splitNum(num int) int {
	numLen := 0
	n := make([]int, 10)
	for i := num; i != 0; i = i / 10 {
		n[i%10]++
		numLen += 1
	}
	numm := make([]int, 2)
	j := 0
	for i := 0; i < 10 && j < numLen; i++ {

		for n[i] > 0 {
			numm[j%2] = numm[j%2]*10 + i
			j++
			n[i]--
		}
	}
	return numm[0] + numm[1]
}
