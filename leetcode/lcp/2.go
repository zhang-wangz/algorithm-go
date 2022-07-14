package main

func fraction(cont []int) []int {
	size := len(cont)
	// 从后往前推
	fz := 1
	fm := cont[size-1]
	for i := size - 2; i >= 0; i-- {
		fz += cont[i] * fm
		g := GCD(fz, fm)
		fz /= g
		fm /= g
		fz, fm = fm, fz
	}
	return []int{fm, fz}
}

func GCD(a, b int) int {
	if a > b {
		a, b = b, a
	}
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}
