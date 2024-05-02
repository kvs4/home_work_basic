package main

import (
	"fmt"
	"strings"
)

func main() {
	var numberLineX, numberLineY int
	fmt.Print("Write the number of horizontal lines: ")
	_, err := fmt.Scan(&numberLineX)
	if err != nil {
		fmt.Println("scanning error")
		return
	}
	fmt.Print("Write the number of vertical lines: ")
	_, err = fmt.Scan(&numberLineY)
	if err != nil {
		fmt.Println("scanning error")
		return
	}

	for i := 0; i < numberLineY; i++ {
		var lineX strings.Builder
		evenUnEven := 0
		if i%2 == 0 {
			evenUnEven = 0
		} else {
			evenUnEven = 1
		}
		for j := 0; j < numberLineX; j++ {
			if j%2 == evenUnEven {
				lineX.WriteString(" ")
			} else {
				lineX.WriteString("#")
			}
		}
		fmt.Println(lineX.String())
	}
}
