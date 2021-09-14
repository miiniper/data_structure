package main

import "fmt"

func main() {
	//s := "angbdaiogndskfgladm"
	//s := "abac"
	//fmt.Println(lengthOfLongestSubstring1(s))

	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	result := BinarySearch(n, 16, 0, len(n))
	fmt.Println(result)
}

//如何使用一次循环遍历两次数组：时间复杂度由O(n^2)降到O(n).
//使用两个下标 left and right
//满足条件1 移动right， 满足条件2 移动left 。
//learn from leetCode 3.
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// BinarySearch 二分法查找实现
func BinarySearch(array []int, target, left, right int) int {
	if len(array) == 0 {
		return 0
	}
	mid := (right + left) / 2
	//	fmt.Println(mid)
	if array[mid] > target {
		return BinarySearch(array, target, left, mid-1)
	} else if array[mid] < target {
		return BinarySearch(array, target, mid+1, right)
	}
	return mid
}
