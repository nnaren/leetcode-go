package y1267

func countServers(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    ans := 0
    for i:=0 ; i<m ; i++ {
        conn := 0
        last := -1
        for j:=0; j<n;j++ {
            if grid[i][j] ==1 {
                conn++
                grid[i][j]++
                last = j
            }
        }
        if conn==1 {
            grid[i][last] = 1
        }
        if conn>1 {
            ans += conn
        }
    }

    for j:=0 ; j<n ;j++ {
        conn:=0
        dup:=0
        for i:=0;i<m;i++ {
            if grid[i][j] == 1{
                conn++
            }
            if grid[i][j] == 2 {
                conn++
                dup++
            }
        }
        if conn>1 {
            ans += (conn-dup)
        }
    }
    return ans
}

