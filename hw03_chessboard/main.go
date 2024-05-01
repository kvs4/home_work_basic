package main

import "fmt"

func main() {
	var numberLineX, numberLineY int
	fmt.Print("Write the number of horizontal lines: ")
	fmt.Scan(&numberLineX)
	fmt.Print("Write the number of vertical lines: ")
	fmt.Scan(&numberLineY)

	for i := 0; i < numberLineY; i++ {
		lineX := ""
		evenUnEven := 0
		if i%2 == 0 {
			evenUnEven = 0
		} else {
			evenUnEven = 1
		}
		for j := 0; j < numberLineX; j++ {
			if j%2 == evenUnEven {
				lineX += " "
			} else {
				lineX += "#"
			}
		}
		fmt.Println(lineX)
	}
}
