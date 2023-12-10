package y946

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(popped)
	var saved []int
	j := 0
label:
	for i := 0; i < n; i++ {
		saveLen := len(saved)
		// 从栈中取
		if saveLen > 0 && saved[saveLen-1] == popped[i] {
			saved = saved[:saveLen-1]
			continue
		}

		// 直接从push取
		for ; j < len(pushed); j++ {
			if pushed[j] == popped[i] {
				j++
				continue label
			}
			saved = append(saved, pushed[j])
		}

		return false
	}
	return true
}
