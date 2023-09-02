package y2682


func circularGameLosers(n int, k int) []int {
    a :=  map[int]int{}

    for i:=0; i<n ;i++ {
        a[i] = 0
    }
    cur := 0
    for i:=0; ;i++{
        cur = (cur+i*k) % n
        a[cur] ++
        if a[cur]==2{
            break
        }
    }
    var ans  []int
    for i:=0;i<n;i++{
        if a[i] == 0 {
            ans = append(ans, i+1)
        }
    }
    return ans
}
