package godcc

import "time"

type Date interface {
	Year() int
	Month() time.Month
	Day() int
}
