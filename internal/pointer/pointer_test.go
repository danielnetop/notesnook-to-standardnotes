package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTo_String(t *testing.T) {
	strings := []string{"a", "b"}

	tests := []struct {
		value    string
		expected *string
	}{
		{value: strings[0], expected: &strings[0]},
		{value: strings[1], expected: &strings[1]},
	}
	for _, test := range tests {
		t.Run("To string", func(t *testing.T) {
			assert.Equal(t, test.expected, To(test.value))
		})
	}
}

func TestTo_Int(t *testing.T) {
	ints := []int{1, 2}

	tests := []struct {
		value    int
		expected *int
	}{
		{value: ints[0], expected: &ints[0]},
		{value: ints[1], expected: &ints[1]},
	}
	for _, test := range tests {
		t.Run("To int", func(t *testing.T) {
			assert.Equal(t, test.expected, To(test.value))
		})
	}
}

func TestTo_Float32(t *testing.T) {
	float32s := []float32{23.4567, -1.1234}

	tests := []struct {
		value    float32
		expected *float32
	}{
		{value: float32s[0], expected: &float32s[0]},
		{value: float32s[1], expected: &float32s[1]},
	}
	for _, test := range tests {
		t.Run("To float32", func(t *testing.T) {
			assert.Equal(t, test.expected, To(test.value))
		})
	}
}

func TestTo_Bool(t *testing.T) {
	bools := []bool{true, false}

	tests := []struct {
		value    bool
		expected *bool
	}{
		{value: bools[0], expected: &bools[0]},
		{value: bools[1], expected: &bools[1]},
	}
	for _, test := range tests {
		t.Run("To bool", func(t *testing.T) {
			assert.Equal(t, test.expected, To(test.value))
		})
	}
}
