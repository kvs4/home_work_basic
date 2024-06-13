package printer

import (
	"testing"

	"github.com/kvs4/home_work_basic/hw06_testing/hw02_fix_app/reader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetStrStaff(t *testing.T) {
	const path string = "data.json"
	const strWant string = `User ID: 10; Age: 25; Name: Rob; Department ID: 3; 
User ID: 11; Age: 30; Name: George; Department ID: 2; `

	staff, err := reader.ReadJSON(path)
	require.NoError(t, err)

	strGot := GetStrStaff(staff)
	assert.Equal(t, strWant, strGot)
}
