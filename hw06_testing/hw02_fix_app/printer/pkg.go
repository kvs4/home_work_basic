package printer

import (
	"fmt"
	"strings"

	"github.com/kvs4/home_work_basic/hw06_testing/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	str := GetStrStaff(staff)
	fmt.Println(str)
}

func GetStrStaff(staff []types.Employee) string {
	var str strings.Builder
	for i := 0; i < len(staff); i++ {
		if i > 0 {
			str.WriteString("\n")
		}
		str.WriteString(fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID))
	}
	return str.String()
}
