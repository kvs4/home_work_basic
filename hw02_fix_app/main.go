package main

import (
	"fmt"
	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func main() {
	path := "data.json"
	var err error

	fmt.Printf("Enter data file path: ")
	_, err = fmt.Scanln(&path)
	if err != nil {
		fmt.Println(err)
		return
	}

	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	printer.PrintStaff(staff)
}
