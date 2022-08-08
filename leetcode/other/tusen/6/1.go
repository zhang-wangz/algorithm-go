package main

// https://leetcode.cn/problems/excel-sheet-column-number/
func titleToNumber(columnTitle string) (number int) {
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		number += int(k) * multiple
		multiple *= 26
	}
	return
}
