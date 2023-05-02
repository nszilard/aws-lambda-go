package pkg

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		c    int
		e    int
	}{
		{
			name: "1st digit",
			c:    1,
			e:    1,
		},
		{
			name: "2nd digit",
			c:    2,
			e:    1,
		},
		{
			name: "3rd digit",
			c:    3,
			e:    2,
		},
		{
			name: "5th digit",
			c:    5,
			e:    5,
		},
		{
			name: "10th digit",
			c:    10,
			e:    55,
		},
	}
	for _, tt := range tests {
		res := Fibonacci(tt.c)

		if res != tt.e {
			t.Errorf("%v should have been: %v, but got: %v", tt.c, tt.e, res)
		}
	}
}
