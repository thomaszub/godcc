package godcc

func base30360(y1, y2, m1, m2, d1, d2 int) float64 {
	return float64(360*(y2-y1)+30*(m2-m1)+d2-d1) / 360.0
}

type D30360BB struct{}

func (_ D30360BB) Frac(from, to Date) float64 {
	d1 := min(from.Day(), 30)
	d2 := to.Day()
	if d1 == 30 {
		d2 = min(d2, 30)
	}
	return base30360(from.Year(), to.Year(), int(from.Month()), int(to.Month()), d1, d2)
}

type D30E360 struct{}

func (_ D30E360) Frac(from, to Date) float64 {
	d1 := min(from.Day(), 30)
	d2 := min(to.Day(), 30)
	return base30360(from.Year(), to.Year(), int(from.Month()), int(to.Month()), d1, d2)
}
