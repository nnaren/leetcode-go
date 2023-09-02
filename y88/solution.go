package y88


// 逆向双指针
func merge(nums1 []int, m int, nums2 []int, n int)  {

    for r1, r2, rr := m-1, n-1, m+n-1 ;  r2>=0; rr--  {
        // 只考虑把nums1的数组移到到末尾情况
        if r1 > -1 &&  nums1[r1] > nums2[r2] {
            nums1[rr] = nums1[r1]
            r1--
        } else { // 其他情况把nums2填充到nums1中
            nums1[rr] = nums2[r2]
            r2--
        }
    }
}


//func merge(nums1 []int, m int, nums2 []int, n int)  {
//
//    for r1, r2, rr := m-1, n-1, m+n-1 ;  r2>=0; rr--  {
//        if r1 == -1 {
//            nums1[rr] = nums2[r2]
//            r2--
//        } else if nums1[r1] > nums2[r2] {
//            nums1[rr] = nums1[r1]
//            r1--
//        } else {
//            nums1[rr] = nums2[r2]
//            r2--
//        }
//    }
//
//}

