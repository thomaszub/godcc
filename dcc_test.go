package godcc

import (
	"testing"
	"time"
)

func date(year int, month time.Month, day int) Date {
	return time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
}

func TestDayCountFractions(t *testing.T) {
	tests := []struct {
		from Date
		to   Date
		exp  float64
	}{
		{
			from: date(2018, 3, 31),
			to:   date(2018, 3, 31),
			exp:  0.0,
		},
		{
			from: date(2018, 4, 30),
			to:   date(2018, 7, 31),
			exp:  0.25,
		},
		{
			from: date(2018, 3, 31),
			to:   date(2018, 9, 30),
			exp:  0.5,
		},
		{
			from: date(2018, 3, 31),
			to:   date(2019, 3, 31),
			exp:  1.0,
		},
	}

	dayCounters := []DayCounter{
		D30360{},
	}

	for _, dc := range dayCounters {
		for _, tt := range tests {
			res := dc.Frac(tt.from, tt.to)
			if res != tt.exp {
				t.Errorf("Wrong result for DayCounter %T from %s to %s. got=%f, exp=%f", dc, tt.from, tt.to, res, tt.exp)
			}
		}
	}

}
