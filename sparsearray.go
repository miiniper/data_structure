//稀疏数组

func main() {
	fmt.Println("    ###############SparseArray########################    ")

	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	ShowArray(chessMap)

	var SparseArray []ValNode

	valnode := ValNode{Col: 11, Row: 11, Val: 0}
	SparseArray = append(SparseArray, valnode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valnode := ValNode{Row: i, Col: j, Val: v2}
				SparseArray = append(SparseArray, valnode)
			}
		}
	}

	fmt.Println("print sparse array:")
	for _, v := range SparseArray {
		fmt.Printf(" %d %d %d\n", v.Row, v.Col, v.Val)
	}

	chessMap2 := Restore(SparseArray)

	ShowArray(chessMap2)
}

func ShowArray(array [11][11]int) {
	for _, v := range array {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)

		}
		fmt.Println()
	}

}

func Restore(SparseArray []ValNode) [11][11]int {

	var chessMap [11][11]int

	for i, j := range SparseArray {
		if i != 0 { // jump first line
			chessMap[j.Row][j.Col] = j.Val
		}

	}

	return chessMap
}
