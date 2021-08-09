//华为笔试，从n楼梯到target,可以*2，+1，-1三种方式上楼梯
package leetcode

func dp(n, target int) int {
	data := make([]int, n+1)
	data[0], data[n] = 0, 1
	i, j := 1, n-1
	for i <= j {
		if i == j {
			data[i] = data[i-1] + 1
			break
		}
		data[i] = data[i-1] + 1
		data[j] = data[j+1] + 1
		i++
		j--
	}
	count := 0
	for target > 2*n {
		if target%2 == 1 {
			count += 2
		} else {
			count++
		}
		target /= 2
	}
	return count + data[target-n]
}

func dfs(n, target int) int {
	data := make([]int, 1)
	data[0] = n
	count := -1
	for {
		l := len(data)
		count++
		for l != 0 {
			if data[0] == target {
				return count
			} else {
				data = append(data, data[0]*2, data[0]+1, data[0]-1)
				data = data[1:]
			}
			l--
		}
	}
}
