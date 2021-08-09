/*
给你一个整数数组 nums，数组中的元素 互不相同。返回该数组所有可能的子集（幂集）。
解集不能包含重复的子集。你可以按 任意顺序 返回解集。
*/
package leetcode

import (
	"math"
)

func subsets(nums []int) [][]int {
	var (
		l   int = len(nums)
		ans     = make([][]int, 0)
		dfs func([]int, int)
	)
	dfs = func(temp []int, index int) {
		for i := index; i < l; i++ {
			//必须使用copy函数复制一个temp切片，否则在回溯的过程中可能出现错误；
			//主要是切片len和cap产生的错误，导致在回溯中改变原切片的值；
			temp1 := make([]int, len(temp)+1)
			copy(temp1, temp)
			temp1 = append(temp1, nums[i])
			dfs(temp1, i+1)
		}
		ans = append(ans, temp)
	}
	dfs([]int{}, 0)
	return ans
}

func subsets2(nums []int) [][]int {
	ans := make([][]int, 1, int(math.Pow(2, float64(len(nums))))+1)
	ans[0] = []int{}
	for _, x := range nums {
		for _, arr := range ans {

			a := make([]int, len(arr), len(arr)+1)
			copy(a, arr)
			ans = append(ans, append(a, x))
		}
	}
	return ans
}
