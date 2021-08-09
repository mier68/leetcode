package leetcode

/*
插入排序:
	每一趟排序，将待排序的关键字插入到已经排好序的一组记录中
分类：
	直接插入排序，折半插入排序，希尔排序
*/

func InsertSort(nums []int) {
	var (
		temp int
		l    int = len(nums)
	)
	for i := 1; i < l; i++ {
		temp = nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if temp < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = temp
	}
}

func BInsertSort(nums []int) {
	var (
		temp int
		l    int = len(nums)
	)
	for i := 1; i < l; i++ {
		temp = nums[i]
		low, high := 0, i-1
		for low <= high {
			m := (low + high) / 2
			if temp < nums[m] {
				high = m - 1
			} else {
				low = m + 1
			}
		}
		for j := i - 1; j >= high+1; j-- {
			nums[j+1] = nums[j]
		}
		nums[high+1] = temp
	}
}

func ShellSort(nums []int, dk int) {
	var (
		temp int
		l    int = len(nums)
	)
	for i := dk; i < l; i++ {
		temp = nums[i]
		j := i - dk
		for ; j >= 0; j -= dk {
			if temp < nums[j] {
				nums[j+dk] = nums[j]
			} else {
				break
			}
		}
		nums[j+dk] = temp
	}
}