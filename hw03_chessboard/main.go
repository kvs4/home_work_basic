package main

import "fmt"

func main() {
	numberLineX := 8
	numberLineY := 8

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
