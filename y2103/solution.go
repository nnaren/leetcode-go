package y2103

import (
	"fmt"
)

func countPoints(rings string) (ans int) {
	l := len(rings)
	m := make(map[byte]byte)
	for i := 0; i < l; i += 2 {
		key := rings[i+1] - '0'
		switch rings[i] {
		case 'B':
			fmt.Printf("%c", rings[i])
			m[key] |= 1
		case 'R':
			fmt.Printf("%c", rings[i])
			m[key] |= 2
		case 'G':
			fmt.Printf("%c", rings[i])
			m[key] |= 4
		}
		//println(rings[i])
	}
	for i := 0; i < 10; i++ {
		v, ok := m[byte(i)]
		if ok && v == 7 {
			ans++
		}
	}
	return
}
