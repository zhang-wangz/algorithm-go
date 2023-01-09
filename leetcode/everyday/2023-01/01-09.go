package main

// 方法1是直接模拟
// g^k(i)≡2^k*i ≡ i (mod(n−1))
// 推理
// 即求 2^k ≡ 1 (mod(n-1)) 最小成立的值
func reinitializePermutation(n int) int {
	if n == 2 {
		return 1
	}
	setp := 1
	pow2 := 2
	for pow2 != 1 {
		setp++
		pow2 = pow2 * 2 % (n -  1)
	}
	return setp
}
