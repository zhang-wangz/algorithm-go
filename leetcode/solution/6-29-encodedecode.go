package solution

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/encode-and-decode-tinyurl/

type Codec struct {
	mp map[string]string
}

func Constructor() Codec {
	return Codec{mp: map[string]string{}}
}

// Encodes a URL to a shortened URL.
func (code *Codec) encode(longUrl string) string {
	strs := strings.Split(longUrl, "/")
	n := len(strs[len(strs)-1])
	boundl := int(math.Pow(float64(10), float64(n)))
	boundr := 0
	for n != 0 {
		boundr = boundr*10 + 9
		n /= 10
	}
	ran := int(boundl + int(float64(boundr-boundl)*rand.Float64()))
	strs[len(strs)-1] = strconv.Itoa(ran)
	code.mp[strings.Join(strs, "/")] = longUrl
	return strings.Join(strs, "/")
}

// Decodes a shortened URL to its original URL.
func (code *Codec) decode(shortUrl string) string {
	if v, ok := code.mp[shortUrl]; ok {
		return v
	} else {
		return ""
	}
}
