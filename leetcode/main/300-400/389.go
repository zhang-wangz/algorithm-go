package main

func findTheDifference(s string, t string) byte {
	num := 0
	for i, c := range s {
		num = num ^ int(c-'a') ^ int(t[i]-'a')
	}
	num ^= int(t[len(t)-1] - 'a')
	return byte(num + 'a')
}
