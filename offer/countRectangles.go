package main

import "fmt"

//(x1,y2)(x2,y2)
//(x1,y1)(x2,y1)

func countRectangles(xy [][]int) int {
	x1y2 := 0
	x2y1 := 0
	count := 0
	for i, _ := range xy {
		x1 := xy[i][0]
		y1 := xy[i][1]

		for _, j := range xy {
			x2 := j[0]
			y2 := j[1]
			if x1 != x2 && x1 != y2 && y1 != x2 && y1 != y2 {
				for k, _ := range xy {
					if xy[k][0] == x1 && xy[k][1] == y2 {
						x1y2++
						for n, _ := range xy {
							if xy[n][0] == x2 && xy[n][1] == y1 {
								x2y1++
								count++
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(x1y2, x2y1)
	return count / 2
}

//[[1,1],[1,3],[3,1],[3,3],[2,2]] = 1
//[[1,1],[1,3], [1,4], [3,1], [3,3], [3,4] ]= 3
//[[1,1], [1,2], [1,3], [2,2], [2,3], [3,1], [3,2] ]=2
func main() {
	a1 := [][]int{{1, 1}, {1, 3}, {1, 4}, {3, 1}, {3, 3}, {3, 4}}
	//a1 := [][]int{{1, 1}, {1, 3}, {3, 1}, {2, 2}, {3, 3}}
	//a1 := [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 2}, {2, 3}, {3, 1}, {3, 2}}
	fmt.Println(countRectangles(a1))
}

func get2(a1 [][]int) {

}
