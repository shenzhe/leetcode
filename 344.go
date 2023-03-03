package code

func Execuate_344(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	sb := []byte(s)
	reverseString(sb)
	return string(sb)
}

func reverseString(s []byte) {
	lo, hi := 0, len(s)-1
	for lo < hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}
