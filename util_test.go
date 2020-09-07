package godcc

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year   int
		isLeap bool
	}{
		{1900, false},
		{2000, true},
		{2001, false},
		{2004, true},
	}
	for _, tt := range tests {
		res := isLeapYear(tt.year)
		if res != tt.isLeap {
			t.Errorf("Wrong result for isLeapYear. Got=%t, exp=%t", res, tt.isLeap)
		}
	}
}
