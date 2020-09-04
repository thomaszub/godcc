package godcc

func min(value1, value2 int) int {
	if value2 < value1 {
		return value2
	}
	return value1
}

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}
