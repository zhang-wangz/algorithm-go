package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct {
	l, r, v int
}

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)

}

func main() {
	// open, _ := os.Open("/Users/zhang-wangz/Desktop/algorithm/tea_0x3f/00_cfinput.txt")
	// solve(open, os.Stdout)
	// defer open.Close()
	solve(os.Stdin, os.Stdout)
}
