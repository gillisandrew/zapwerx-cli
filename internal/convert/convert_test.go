package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInteger(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected int
		err      bool
	}{
		{
			name:     "valid int",
			input:    1,
			expected: 1,
			err:      false,
		},
		{
			name:     "valid string",
			input:    "1",
			expected: 1,
			err:      false,
		},
		{
			name:  "invalid string",
			input: "1asdaw",
			err:   true,
		},
		{
			name:  "invalid bool",
			input: false,
			err:   true,
		},
	}

	for _, tc := range testCases {
		res, err := ToInteger(tc.input)

		if tc.err {
			if assert.Error(t, err) {
				assert.Equal(t, errConversionError(tc.input), err)
			}
		} else {
			assert.Equal(t, tc.expected, res)
		}
	}
}
