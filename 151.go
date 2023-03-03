package code

import "fmt"

func Execuate_151(s string) string {
	return reverseWords(s)
}

func reverseWords(sb string) string {
	// sb := []byte(s)
	l := len(sb)
	// sb[0], sb[l-1] = sb[l-1], sb[0]
	//the sky is blue
	lo, hi := trimSpace(sb, 0, l-1)
	res := make([]byte, 0)
	fast := hi
	for hi >= lo {
		fmt.Printf("fast=%d, hi=%d, lo=%d\n", fast, hi, lo)
		if fast >= lo && sb[fast] != ' ' {
			fast--
		} else if fast == hi && hi == lo {
			res = append(res, sb[lo])
			break
		} else {
			res = append(res, sb[fast+1:hi+1]...)
			if fast < lo {
				break
			}
			hi = fast - 1

			//去掉多有空格
			for sb[hi] == ' ' {
				hi--
				if hi < lo {
					break
				}
			}
			//加入一个空格
			res = append(res, ' ')
			fast = hi
		}

	}

	return string(res)
}

func trimSpace(sb string, lo, hi int) (int, int) {
	//去掉首空格
	for lo < hi {
		if sb[lo] == ' ' {
			lo++
		} else {
			break
		}
	}
	//去掉尾空格
	for hi > 0 {
		if sb[hi] == ' ' {
			hi--
		} else {
			break
		}
	}
	return lo, hi
}
