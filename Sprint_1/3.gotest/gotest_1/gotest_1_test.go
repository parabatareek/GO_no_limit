package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "simple test #1",
			value: -3,
			want:  3,
		},
		{
			name:  "simple test #2",
			value: 3,
			want:  3,
		},
		{
			name:  "simple test #3",
			value: -2.000001,
			want:  2.000001,
		},
		{
			name:  "simple test #4",
			value: -0.000000003,
			want:  -0.000000003,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if abs := Abs(test.value); abs != test.want {
				t.Errorf("Abs() result = %v, want %v", abs, test.want)
			}
		})
	}
}
