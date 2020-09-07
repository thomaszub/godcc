package godcc

import (
	"testing"
	"time"
)

type DcTest struct {
	from   Date
	to     Date
	dcExps []DcExp
}

type DcExp struct {
	dcc DayCounter
	exp float64
}

func date(year int, month time.Month, day int) Date {
	return time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
}

func TestDayCountFractions(t *testing.T) {
	tests := []DcTest{
		{
			from: date(2018, 3, 31),
			to:   date(2018, 3, 31),
			dcExps: []DcExp{
				{
					dcc: D30360BB{},
					exp: 0,
				},
				{
					dcc: D30E360{},
					exp: 0,
				},
			},
		},
		{
			from: date(2018, 4, 30),
			to:   date(2018, 7, 31),
			dcExps: []DcExp{
				{
					dcc: D30360BB{},
					exp: 0.25,
				},
				{
					dcc: D30E360{},
					exp: 0.25,
				},
			},
		},
		{
			from: date(2018, 3, 31),
			to:   date(2018, 9, 30),
			dcExps: []DcExp{
				{
					dcc: D30360BB{},
					exp: 0.5,
				},
				{
					dcc: D30E360{},
					exp: 0.5,
				},
			},
		},
		{
			from: date(2018, 3, 31),
			to:   date(2019, 3, 31),
			dcExps: []DcExp{
				{
					dcc: D30360BB{},
					exp: 1.0,
				},
				{
					dcc: D30E360{},
					exp: 1.0,
				},
			},
		},
	}

	for _, tt := range tests {
		for _, dct := range tt.dcExps {
			res := dct.dcc.Frac(tt.from, tt.to)
			t.Logf("Testing DayCounter %T from %s to %s", dct.dcc, tt.from, tt.to)
			if res != dct.exp {
				t.Errorf("Wrong result for DayCounter %T from %s to %s. got=%f, exp=%f", dct.dcc, tt.from, tt.to, res, dct.exp)
			}
		}
	}

}
