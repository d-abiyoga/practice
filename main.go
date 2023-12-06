package main

func Pow(a, b float32) float32 {
	ret := float32(1)
	isNegative := false
	if b < 0 {
		b = -b
		isNegative = true
	}

	isBEven := int(b)%2 == 0

	for i := float32(0); i < b/2; i++ {
		ret = ret * a
	}

	if isBEven {
		ret = ret * ret
	} else {
		ret = ret * ret * a * float32(int(b)%2)
	}

	if isNegative {
		ret = 1 / ret
	}

	return ret
}
