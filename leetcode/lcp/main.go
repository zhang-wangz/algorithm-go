package main

func game(guess []int, answer []int) int {
	cnt := 0
	for i := range guess {
		if guess[i] == answer[i] {
			cnt++
		}
	}
	return cnt
}
