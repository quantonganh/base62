package base62

import (
	"math"
	"strings"
)

const (
	base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b = 62
)

func Encode(n int) string {
	var s string
	for n != 0 {
		remainder := n % b
		n = int(math.Floor(float64(n / b)))
		s = string(base[remainder]) + s
	}
	return s
}

func Decode(s string) int {
	n := 0
	for _, r := range s {
		n = (b * n) + strings.Index(base, string(r))
	}
	return n
}