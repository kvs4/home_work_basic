package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringEmployee(t *testing.T) {
	const strWant string = "User ID: 10; Age: 25; Name: Rob; Department ID: 3; "
	empl := Employee{
		UserID:       10,
		Age:          25,
		Name:         "Rob",
		DepartmentID: 3,
	}

	strGot := empl.String()
	assert.Equal(t, strWant, strGot)
}

func TestStringEmployeeEmpty(t *testing.T) {
	const strWant string = "User ID: 0; Age: 0; Name: ; Department ID: 0; "
	empl := Employee{
		UserID:       0,
		Age:          0,
		Name:         "",
		DepartmentID: 0,
	}

	strGot := empl.String()
	assert.Equal(t, strWant, strGot)
}
