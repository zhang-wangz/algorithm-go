package main

func winnerSquareGame(n int) bool {
	f := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		for k := 1; k*k <= i; k++ {
			if !f[i-k*k] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}
