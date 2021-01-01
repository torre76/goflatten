package flatten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlattenInt(t *testing.T) {
	res, err := Int(
		[]interface{}{
			[]interface{}{1, 2,
				[]int{3}},
			[]interface{}{4, 5,
				[]interface{}{6,
					[]int{7, 8}}}})

	assert.NoError(t, err)
	assert.Equal(t, res, []int{1, 2, 3, 4, 5, 6, 7, 8})

	res, err = Int(
		[]interface{}{
			[]interface{}{1, "2",
				[]int{3}},
			[]interface{}{4, 5,
				[]interface{}{6,
					[]int{7, 8}}}})

	assert.Error(t, err)

	res, err = Int("Ciao")

	assert.Error(t, err)
}
