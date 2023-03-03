package code

import (
	"strconv"
	"strings"
	"unicode"
)

type Calculator struct {
	origin  string
	s       string
	left    int
	signMap map[byte]bool
}

var CalculatorErrorCode = map[string]int{
	"sign_error":    1,
	"divide0":       2,
	"more(":         3,
	"more)":         4,
	"empty_error":   5,
	"express_error": 6,
}

func NewCalculator(s string) Calculator {
	signMap := map[byte]bool{
		'+': true,
		'-': true,
		'*': true,
		'/': true,
		'(': true,
		')': true,
		' ': true,
	}
	return Calculator{
		origin:  s,
		s:       strings.TrimSpace(s),
		left:    0,
		signMap: signMap,
	}
}

func (cal *Calculator) setErrorCode(s string) string {
	if code, ok := CalculatorErrorCode[s]; ok {
		return "[" + strconv.Itoa(code) + "]"
	}
	return "[-1]"
}

func (cal *Calculator) getPanicCode(s string) int {
	l := len(s)
	if l >= 3 {
		c := s[1]
		if unicode.IsDigit(rune(c)) {
			return int(c) - int('0')
		}
	}
	return -1
}

func (cal *Calculator) CheckPanicErrorCode(s string, code int) bool {
	errCode := cal.getPanicCode(s)
	// fmt.Println("check error", errCode, code)
	return errCode == code
}

func (cal *Calculator) Execuate() int {
	stack := NewStack()
	num := 0
	numFlag := false
	emptyFlag := false
	sign := '+'
	//往前遍历字符串
	for len(cal.s) > 0 {
		c := cal.s[0]
		cal.s = cal.s[1:]
		// fmt.Println(string(c), cal.s)
		//空字符串直接跳过，除了最后一个空格
		if c == ' ' {
			emptyFlag = true
			continue
		}
		//判断是不是数字，考虑到连续的数字进行累加
		if unicode.IsDigit(rune(c)) {
			if numFlag {
				//数字之间不能有空格
				if emptyFlag {
					panic(cal.setErrorCode("empty_error") + cal.origin + ": express error, empty can not between in numbers")
				}
			} else {
				numFlag = true
			}
			num = num*10 + (int(c) - int('0'))
		} else {
			numFlag = false
			if _, ok := cal.signMap[c]; !ok {
				panic(cal.setErrorCode("sign_error") + cal.origin + ": express error, invaild sign:" + string(c))
			}
		}
		emptyFlag = false
		//如果是(，递归处理，直到 碰到）弹出，结果返回
		if c == '(' {
			cal.left++
			num = cal.Execuate()
		}
		//是操作符，或者到了字符串最后，进行sush操作
		if len(cal.s) == 0 || !unicode.IsDigit(rune(c)) {
			switch sign {
			case '+': //直接入stack
				stack.Push(num)
			case '-': //变成负数，再入stack
				num = 0 - num
				stack.Push(num)
			case '*': //取出statck的top元素并弹出，相乘，再入stack
				v, _ := stack.Pop()
				stack.Push(num * v)
			case '/': //取出statck的top元素并弹出，相乘，再入stack
				if num == 0 {
					panic(cal.setErrorCode("divide0") + cal.origin + ":" + cal.s + ": can not divide zero")
				}
				v, _ := stack.Pop()
				stack.Push(v / num)
			}

			sign = rune(c)
			num = 0
			numFlag = false
		}

		if c == ')' {
			cal.left--
			if cal.left < 0 {
				panic(cal.setErrorCode("more)") + cal.origin + ": express error, more )")
			}
			break
		}
	}

	if len(cal.s) == 0 && cal.left > 0 {
		panic(cal.setErrorCode("more(") + cal.origin + ":" + cal.s + ": express error, more (: " + strconv.Itoa(cal.left))
	}
	//statck求和，返回
	return stack.Sum()
}
