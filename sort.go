package main

import (
	"fmt"
)

func main() {
	fmt.Println("##################sort##########################")
	aa := []int{1, 3, 321, 45, 788, 12, 38, 425, 436, 233, 88, 32, 9, 65, 25, 27, 56}
	fmt.Println("初始化：    ", aa)
	fmt.Println("倒序：      ", SelectSort(aa))
	//	fmt.Println("再查看aaa   ", aa)
	fmt.Println("正序：      ", SelectSort2(aa))
	//	fmt.Println("再查看aaa   ", aa)

}

//倒叙:选择排序
func SelectSort(array []int) []int {

	for k := 0; k < len(array); k++ {
		max := array[k]
		index := k

		//循环 len-1次，内部循环找出最大值，并和k交换
		for i := k + 1; i < len(array); i++ {
			if max < array[i] {
				max = array[i]
				index = i
			}
		}
		if index != k {
			array[k], array[index] = array[index], array[k]
		}
	}

	return array
}

//正序
func SelectSort2(array []int) []int {
	aa := array

	for k := 0; k < len(aa); k++ {
		min := aa[k]
		index := k

		for i := k + 1; i < len(aa); i++ {
			if min > aa[i] {
				min = aa[i]
				index = i
			}
		}

		if index != k {
			aa[k], aa[index] = aa[index], aa[k]
		}
	}

	return aa
}
