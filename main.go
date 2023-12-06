package main

func Pow(a, b float32) float32 {
	ret := float32(1)
	isNegative := false
	if b < 0 {
		b = -b
		isNegative = true
	}

	for i := float32(0); i < b; {
		if i == float32(0) {
			ret = a
			i++
			continue
		}

		if i*2 < b {
			ret = ret * ret
			i *= 2
			continue
		}

		ret = ret * a
		i++
	}

	if isNegative {
		ret = 1 / ret
	}

	return ret
}
