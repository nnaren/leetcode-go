package y207CourseSchedule

import (
	"testing"
)

func Test_solv(t *testing.T) {

	numCourses := 2
	prerequisites :=  [][]int{{1,0},{0,1}}
	canFinish2(numCourses, prerequisites)
}
