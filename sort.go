package main

import (
	"fmt"
)

func main() {
	fmt.Println("##################sort##########################")
	//aa := []int{30, 61, 321, 45, 788, 12, 38, 425, 436, 9876, 233, 88, 32, 9, 65, 25, 27, 56}
	//aa := []int{3, 6, 9, 2, 5, 8, 1, 4, 7}
	aa := []int{3, 1, 4, 5, 2}
	fmt.Println("init array == ", aa)
	//fmt.Println("select倒序：      ", SelectSort(aa))
	//	fmt.Println("再查看aaa   ", aa)
	//fmt.Println("select正序：      ", SelectSort2(aa))
	//	fmt.Println("再查看aaa   ", aa)
	//InsertSort(aa)
	//InsertSort2(aa)
	//fmt.Println(aa)
	//QuickSort(aa, 0, len(aa)-1)
	//fmt.Println(aa)
	//fmt.Println(min(aa))
	//fmt.Println(max(aa))
	//bubbleSort(aa)

}

//取左边为基准，除去left的数组分为左右
//快速排序，正序
func QuickSort(array []int, left, right int) {
	//	a := array
	if left >= right {
		return
	}
	explodeIndex := left
	for i := left + 1; i <= right; i++ {
		if array[left] >= array[i] {
			explodeIndex++
			array[i], array[explodeIndex] = array[explodeIndex], array[i]
		}
		fmt.Println(array)
	}

	array[left], array[explodeIndex] = array[explodeIndex], array[left]

	//fmt.Println("***************************进入左递归*****************************")
	QuickSort(array, left, explodeIndex-1)
	//fmt.Println("###########################进入右递归#############################")
	QuickSort(array, explodeIndex+1, right)
}

//快速排序，倒序
func QuickSort2(array []int, left, right int) {
	//	a := array
	if left >= right {
		return
	}
	explodeIndex := left
	for i := left + 1; i <= right; i++ {
		if array[left] <= array[i] {
			explodeIndex++
			array[i], array[explodeIndex] = array[explodeIndex], array[i]
		}
		fmt.Println(array)
	}
	array[left], array[explodeIndex] = array[explodeIndex], array[left]

	//fmt.Println("***************************进入左递归*****************************")
	QuickSort2(array, left, explodeIndex-1)
	//fmt.Println("###########################进入右递归#############################")
	QuickSort2(array, explodeIndex+1, right)
}

//插入，正序
func InsertSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] { //当前元素和前一个比较，小于就交换位置
				array[j], array[j-1] = array[j-1], array[j]
			} else {
				break
			}
		}
		//fmt.Println(array)
	}
	fmt.Println(array)
}

//插入，倒序
func InsertSort2(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] > array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			} else {
				break
			}
		}
	}

	fmt.Println(array)

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

//最小值
func min(arr []int) int {
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[index] > arr[i] {
			index = i
		}
	}
	return arr[index]
}

func max(arr []int) int {
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[index] < arr[i] {
			index = i
		}
	}
	return arr[index]
}

//冒泡排序,不优
func bubbleSort(arr []int) {
	for k := 0; k < len(arr)-1; k++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
			fmt.Println(arr)
		}
	}
	fmt.Println(arr)
}
