package time

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMilliToTime(t *testing.T) {
	tests := []struct {
		milli int64
		want  string
	}{
		{
			milli: 1699404581000,
			want:  "2023-11-08 00:49:41Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			assert.Equal(t, tt.want, MilliToTime(tt.milli))
		})
	}
}
