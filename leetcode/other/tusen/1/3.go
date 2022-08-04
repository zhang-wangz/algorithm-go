package main

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type node struct {
	s string
	v int
}

func countOfAtoms(formula string) string {
	n := len(formula)
	str := []byte(formula)
	m := map[string]int{}
	st := make([]string, 0)
	idx := 0
	for i := 0; i < n; {
		c := str[i]
		if c == '(' || c == ')' {
			st = append(st, string(c))
			i++
		} else {
			if unicode.IsDigit(rune(c)) {
				j := i
				for j < n && unicode.IsDigit(rune(str[j])) {
					j++
				}
				numstr := str[i:j]
				i = j
				cnt, _ := strconv.Atoi(string(numstr))

				if len(st) != 0 && st[len(st)-1] == ")" {
					tmp := make([]string, 0)
					// pop )
					st = st[:len(st)-1]
					for len(st) != 0 && st[len(st)-1] != "(" {
						cur := st[len(st)-1]
						st = st[:len(st)-1]
						if _, ok := m[cur]; !ok {
							m[cur] = 1
						}
						m[cur] = m[cur] * cnt
						tmp = append(tmp, cur)
					}
					st = st[:len(st)-1] // pop (

					for k := len(tmp) - 1; k >= 0; k-- {
						st = append(st, tmp[k])
					}

					// 如果栈顶元素不是 )，说明当前数值只能应用给栈顶的原子
				} else {
					cur := st[len(st)-1]
					st = st[:len(st)-1]
					if _, ok := m[cur]; !ok {
						m[cur] = 1
					}
					m[cur] *= cnt
					st = append(st, cur)
				}
			} else {
				// 获取完整的原子名
				j := i + 1
				for j < n && unicode.IsLower(rune(str[j])) {
					j++
				}
				cur := string(str[i:j]) + "_" + strconv.Itoa(idx)
				idx++
				m[cur] += 1
				i = j
				st = append(st, cur)
			}
		}
	}
	// 将不同编号的相同原子进行合并
	mm := map[string]node{}
	for k, v := range m {
		atom := strings.Split(k, "_")[0]
		var nd node
		if _, ok := mm[atom]; ok {
			nd = mm[atom]
		} else {
			nd = node{atom, 0}
		}
		nd.v += v
		mm[atom] = nd
	}
	q := make([]node, 0)
	for _, v := range mm {
		q = append(q, v)
	}
	sort.Slice(q, func(i, j int) bool {
		return q[i].s < q[j].s
	})
	sb := strings.Builder{}
	for len(q) != 0 {
		poll := q[0]
		q = q[1:]
		sb.WriteString(poll.s)
		if poll.v > 1 {
			sb.WriteString(strconv.Itoa(poll.v))
		}
	}
	return sb.String()
}
