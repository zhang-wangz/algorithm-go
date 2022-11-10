package _1_60

func isAlienSorted(words []string, order string) bool {
	morder := map[byte]int{}
	for i := 0; i < len(order); i++ {
		morder[order[i]] = i
	}
out:
	for i := 0; i < len(words)-1; i++ {
		w0 := words[i]
		w1 := words[i+1]
		for j := 0; j < len(w0); j++ {
			if j < len(w1) && morder[w0[j]] > morder[w1[j]] {
				return false
			} else if j < len(w1) && morder[w0[j]] < morder[w1[j]] {
				continue out
			} else if j < len(w1) && morder[w0[j]] == morder[w1[j]] {
				continue
			}
			if j >= len(w1) {
				return false
			}
		}
	}
	return true
}
