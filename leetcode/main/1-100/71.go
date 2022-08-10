package main

import "strings"

func simplifyPath(path string) string {
	stk := make([]string, 0)
	data := strings.Split(path, "/")
	for _, name := range data {
		if name == ".." {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}
		} else if name != "" && name != "." {
			stk = append(stk, name)
		}
	}
	return "/" + strings.Join(stk, "/")
}
