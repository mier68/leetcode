package leetcode

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	data := []int{3, 2, 1}
	result := []int{1, 2, 3}
	InsertSort(data)
	for i := 0; i < 3; i++ {
		if result[i] != data[i] {
			t.Errorf("InsertSort:want:%v,result:%v", result, data)
			break
		}
	}
}

func TestBInsertSort(t *testing.T) {
	data := []int{3, 2, 1}
	result := []int{1, 2, 3}
	BInsertSort(data)
	for i := 0; i < 3; i++ {
		if result[i] != data[i] {
			t.Errorf("BInsertSort:want:%v,result:%v", result, data)
			break
		}
	}
}

func TestBShellSort(t *testing.T) {
	data := []int{49, 38, 65, 97, 76, 13, 27, 49, 55, 4}
	result := []int{4, 13, 27, 38, 49, 49, 55, 65, 76, 97}
	dk := []int{5, 3, 1}
	for i := 0; i < len(dk); i++ {
		ShellSort(data, dk[i])
	}
	for i := 0; i < 3; i++ {
		if result[i] != data[i] {
			t.Errorf("ShellSort:want:%v,result:%v", result, data)
			break
		}
	}
}
