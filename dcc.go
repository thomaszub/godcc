package godcc

type DayCounter interface {
	Frac(from, to Date) float64
}
