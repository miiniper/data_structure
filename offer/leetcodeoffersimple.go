package main

import (
	"fmt"
	"strings"
)

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

//股票最大收益
func maxProfit(prices []int) int {
	//tmp := prices[0]
	res := 0
	for j := 0; j < len(prices); j++ {
		for i := j + 1; i < len(prices); i++ {
			r := prices[i] - prices[j]
			if res < r {
				res, r = r, res
			}
		}

	}

	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//不含重复字符子串的最大长度
//abcabcbb
func lengthOfLongestSubstring(s string) int {
	left := 0
	le := 0

	m1 := make(map[uint8]int)

	for right := 0; right < len(s); right++ {
		if j, ok := m1[s[right]]; ok {
			left = max(left, j)
		}

		m1[s[right]] = right + 1
		le = max(right-left+1, le)
	}
	return le
}

//三数之和
func threeSum(arr []int) [][3]int {
	res := [][3]int{}
	for i := 0; i < len(arr); i = i + 3 {
		for j := i + 1; j < len(arr); j = j + 2 {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					res = append(res, [3]int{arr[i], arr[j], arr[k]})
					//fmt.Println(res)
				}
				//fmt.Println(i, j, k)
			}
		}
	}
	//res =res[][:2]
	return removeThreeSum(res)
}

func removeThreeSum(arr [][3]int) [][3]int {
	res := [][3]int{}
	m1 := make(map[[3]int]int)

	for i := 0; i < len(arr); i++ {
		if _, ok := m1[arr[i]]; ok {
			continue
		}
		res = append(res, arr[i])
		m1[arr[i]] = i + 1
	}
	return res

}

//数组去重
func removeNumArray(arr []int) []int {
	var res []int
	m1 := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := m1[arr[i]]; ok {
			continue
		}
		res = append(res, arr[i])
		m1[arr[i]] = i + 1
	}
	return res
}

func insertSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			} else {
				break
			}
		}
	}
	//	fmt.Println(array)
	return array
}

//一下是测试用的
//main is demo test
func main() {
	a := []int{1, 3, 4, 4, 6, 8, 9, 6, 5, 4, 3, 0}
	//fmt.Println(findRepeatNumber(a))
	//arr2 := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	//findNumberIn2DArray(arr2, 5)
	//s := "We are happy."
	//fmt.Println(replaceSpace(s))
	//fmt.Println(maxProfit(a))
	//s := "pwwkew"
	//a := []int{-1, 0, 1, 2, -1, -4}
	//a := []int{0, 0, 0, 0, 0}
	fmt.Println(insertSort(a))

}
