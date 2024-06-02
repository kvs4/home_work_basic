package hw03

import (
	"fmt"
	"strings"
)

func Chessboard() {
	var numberLineX, numberLineY int
	var res string

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

	res, err = GetStrChessboard(numberLineX, numberLineY)
	if err != nil {
		fmt.Println("printing chessboard error")
		return
	}

	fmt.Println(res)
}

func GetStrChessboard(numberLineX, numberLineY int) (string, error) {
	var resStr strings.Builder
	for i := 0; i < numberLineY; i++ {
		if i > 0 {
			resStr.WriteString("\n")
		}
		evenUnEven := 0
		if i%2 == 0 {
			evenUnEven = 0
		} else {
			evenUnEven = 1
		}
		for j := 0; j < numberLineX; j++ {
			if j%2 == evenUnEven {
				resStr.WriteString(" ")
			} else {
				resStr.WriteString("#")
			}
		}
	}
	return resStr.String(), nil
}
