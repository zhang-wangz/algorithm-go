package main

// 官解
func findKthBit(n int, k int) byte {
	invert := func(bit byte) byte {
		return '0' + '1' - bit
	}
	if k == 1 {
		return '0'
	}
	mid := 1 << (n - 1)
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		// 在后半部分寻找第k位
		// 如果k > mid， 求其对应的在前半部分的映射位
		k = mid*2 - k
		return invert(findKthBit(n-1, k))
	}
}

func findKthBit2(n int, k int) byte {
	reverse := func(s string) string {
		i, j := 0, len(s)-1
		b := []byte(s)
		for i < j {
			b[i], b[j] = b[j], b[i]
			if b[i] == '0' {
				b[i] = '1'
			} else {
				b[i] = '0'
			}
			if b[j] == '0' {
				b[j] = '1'
			} else {
				b[j] = '0'
			}
			i++
			j--
		}
		if i == j {
			if b[i] == '0' {
				b[i] = '1'
			} else {
				b[i] = '0'
			}
		}
		return string(b)
	}
	var dfs func(int, string) string
	dfs = func(idx int, s string) string {
		if idx == n {
			return s
		}
		return dfs(idx+1, s+"1"+reverse(s))
	}
	ans := dfs(0, "0")
	return ans[k-1]
}
