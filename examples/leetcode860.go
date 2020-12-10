package main

import "fmt"

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, i := range bills {
		if i == 5 {
			five++
		}
		if i == 10 {
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		}
		if i == 20 {
			if ten > 0 && five > 0 {
				five--
				ten--
			} else if five > 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	a := []int{5, 5, 10, 10, 5, 5, 20}
	fmt.Println(lemonadeChange(a))
}
