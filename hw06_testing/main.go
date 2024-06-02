package main

import (
	"fmt"

	hw03 "github.com/kvs4/home_work_basic/hw06_testing/hw03_chessboard"
	hw04 "github.com/kvs4/home_work_basic/hw06_testing/hw04_struct_comparator"
	hw05 "github.com/kvs4/home_work_basic/hw06_testing/hw05_shapes"
)

func main() {
	hw03.Chessboard()
	fmt.Printf("\n")
	hw04.CompareBook()
	fmt.Printf("\n")
	hw05.PrintCalculateAreaSomeShapes()
}
