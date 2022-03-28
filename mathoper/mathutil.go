package mathoper

func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
