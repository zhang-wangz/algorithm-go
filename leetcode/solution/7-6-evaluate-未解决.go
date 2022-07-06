package solution

import (
	"strconv"
	"unicode"
)

// 状态机
const (
	VALUE = iota // 初始状态
	NONE         // 表达式类型未知
	LET          // let 表达式
	LET1         // let 表达式已经解析了 vi 变量
	LET2         // let 表达式已经解析了最后一个表达式 expr
	ADD          // add 表达式
	ADD1         // add 表达式已经解析了 e1 表达式
	ADD2         // add 表达式已经解析了 e2 表达式
	MULT         // mult 表达式
	MULT1        // mult 表达式已经解析了 e1 表达式
	MULT2        // mult 表达式已经解析了 e2 表达式
	DONE         // 解析完成
)

type expr struct {
	status int
	vr     string // let 的变量 vi
	value  int    // VALUE 状态的数值，或者 LET2 状态最后一个表达式的数值
	e1, e2 int    // add 或 mult 表达式的两个表达式 e1 和 e2 的数值
}

func evaluate(expression string) int {
	scope := map[string][]int{}
	calculateToken := func(token string) int {
		if unicode.IsLower(rune(token[0])) {
			vals := scope[token]
			return vals[len(vals)-1]
		}
		val, _ := strconv.Atoi(token)
		return val
	}

	vars := [][]string{}
	s := []expr{}
	cur := expr{status: VALUE}
	for i, n := 0, len(expression); i < n; {
		if expression[i] == ' ' {
			i++ // 去掉空格
			continue
		}
		if expression[i] == '(' {
			i++ // 去掉左括号
			s = append(s, cur)
			cur.status = NONE
			continue
		}
		var token string
		if expression[i] == ')' { // 本质上是把表达式转成一个 token
			i++ // 去掉右括号
			if cur.status == LET2 {
				token = strconv.Itoa(cur.value)
				for _, vr := range vars[len(vars)-1] { // 清除作用域
					scope[vr] = scope[vr][:len(scope[vr])-1]
				}
				vars = vars[:len(vars)-1]
			} else if cur.status == ADD2 {
				token = strconv.Itoa(cur.e1 + cur.e2)
			} else {
				token = strconv.Itoa(cur.e1 * cur.e2)
			}
			cur, s = s[len(s)-1], s[:len(s)-1] // 获取上层状态
		} else {
			i0 := i
			for i < n && expression[i] != ' ' && expression[i] != ')' {
				i++
			}
			token = expression[i0:i]
		}

		switch cur.status {
		case VALUE:
			cur.value, _ = strconv.Atoi(token)
			cur.status = DONE
		case NONE:
			if token == "let" {
				cur.status = LET
				vars = append(vars, nil) // 记录该层作用域的所有变量, 方便后续的清除
			} else if token == "add" {
				cur.status = ADD
			} else if token == "mult" {
				cur.status = MULT
			}
		case LET:
			if expression[i] == ')' { // let 表达式的最后一个 expr 表达式
				cur.value = calculateToken(token)
				cur.status = LET2
			} else {
				cur.vr = token
				vars[len(vars)-1] = append(vars[len(vars)-1], token) // 记录该层作用域的所有变量, 方便后续的清除
				cur.status = LET1
			}
		case LET1:
			scope[cur.vr] = append(scope[cur.vr], calculateToken(token))
			cur.status = LET
		case ADD:
			cur.e1 = calculateToken(token)
			cur.status = ADD1
		case ADD1:
			cur.e2 = calculateToken(token)
			cur.status = ADD2
		case MULT:
			cur.e1 = calculateToken(token)
			cur.status = MULT1
		case MULT1:
			cur.e2 = calculateToken(token)
			cur.status = MULT2
		}
	}
	return cur.value
}
