package _022_07

import "sort"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	l := make([]int, 6)
	l[0] = countdis(p1, p2)
	l[1] = countdis(p1, p3)
	l[2] = countdis(p1, p4)
	l[3] = countdis(p2, p3)
	l[4] = countdis(p2, p4)
	l[5] = countdis(p3, p4)
	sort.Ints(l)
	return l[0] != 0 && l[0] == l[3] && l[4] == l[5] && l[0]*2 == l[5]
}

func countdis(p1, p2 []int) int {
	e1 := p1[0] - p2[0]
	e2 := p1[1] - p2[1]
	return e1*e1 + e2*e2
}
