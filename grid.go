package leetcode

func numIslands(grid [][]byte) int {
	var (
		l      int = len(grid)
		k      int = len(grid[0])
		dfs    func(int, int)
		InArea func(int, int) bool
		ans    int
	)
	InArea = func(i, j int) bool {
		if i >= 0 && i <= l-1 && j >= 0 && j <= k {
			return true
		}
		return false
	}

	dfs = func(i, j int) {
		if !InArea(i, j) {
			return
		}
		if grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for k1, v1 := range grid {
		for k2, v2 := range v1 {
			if v2 == '1' {
				ans++
				dfs(k1, k2)
			}
		}
	}
	return ans
}
