package solution

import "fmt"

func ReverseString(s []byte) {
	res := make([]byte, len(s))
	str := string(s)
	reverse(str, 0, res)
	fmt.Println(string(res))
}

func reverse(s string, i int, res []byte) {
	if i == len(s) {
		return
	}
	reverse(s, i+1, res)
	res = append(res, s[i])
	//res[len(s)-i-1] = s[i]
}
