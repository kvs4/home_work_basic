package reader

import (
	"testing"

	"github.com/kvs4/home_work_basic/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadJSON(t *testing.T) {
	const path string = "data.json"

	staffWant := []types.Employee{
		{
			UserID:       10,
			Age:          25,
			Name:         "Rob",
			DepartmentID: 3,
		},
		{
			UserID:       11,
			Age:          30,
			Name:         "George",
			DepartmentID: 2,
		},
	}

	staffGot, err := ReadJSON(path)
	require.NoError(t, err)
	assert.Equal(t, staffWant, staffGot)
}

func TestReadJSONEmpty(t *testing.T) {
	const path string = "data_empty.json"

	staffWant := []types.Employee{}

	staffGot, err := ReadJSON(path)
	require.NoError(t, err)
	assert.Equal(t, staffWant, staffGot)
}
