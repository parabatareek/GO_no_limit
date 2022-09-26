package assert_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{name: "simple test #1", value: -3.001, want: 3.001},
		{name: "simple test #2", value: -0.00000001, want: 0.00000001},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := Abs(test.value)
			assert.Equal(t, test.want, v)
		})
	}
}
