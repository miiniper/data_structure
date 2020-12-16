package main

import (
	"fmt"
	"strings"
)

func main() {
	//a := []int{1, 3, 4, 4, 6, 8, 9, 6, 5, 4, 3, 0}
	//fmt.Println(findRepeatNumber(a))
	//arr2 := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	//findNumberIn2DArray(arr2, 5)
	//s := "We are happy."
	//fmt.Println(replaceSpace(s))
	fmt.Println(fib2(6))
}

//给一个数组找出重复数字
func findRepeatNumber(nums []int) []int {
	m := make(map[int]int)
	for _, i := range nums {
		if _, ok := m[i]; ok {
			m[i] = m[i] + 1
		} else {
			m[i] = 1
		}
	}
	fmt.Println(m)
	res := []int{}
	for k, v := range m {
		if v >= 2 {
			res = append(res, k)
		}
	}
	return res
}

//二维数组查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, arr := range matrix {
		for _, i := range arr {
			fmt.Println(i)
			if i == target {
				return true
			}
		}
	}
	return false
}

//字符串替换：空格替换为%20
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

//递归fibonacci数列
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//循环fibonacci数列
func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var f1 int = 0
	var f2 int = 1
	var sum int = 0
	for i := 0; i < n-1; i++ {
		sum = f1 + f2
		f1 = f2
		f2 = sum
	}
	return sum % 1000000007
}
