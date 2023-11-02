package y210CourseSchedule2
import (
	"testing"
)

func Test_solv(t *testing.T) {

	numCourses := 4
	prerequisites :=  [][]int{{1,0},{2,0},{3,1},{3,2}}
	findOrder(numCourses, prerequisites)
}
