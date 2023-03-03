package code

import (
	"fmt"
	"sync"
	"testing"
)

func TestCalculator(t *testing.T) {

	smap := map[string]int{
		"3-1+2":                             4,
		"5-8-3":                             -6,
		"101-88+89-47+45":                   100,
		"3 + (6 - 3) * 20 * 30 / 15":        123,
		"3 + (6 - 3) * (20-3) * 30 / 15":    105,
		" (3 + (6 - 3) * (20-3)) * 30 / 15": 108,
		"( 1+( 4+ 55+2)-(30+2)*32)+(6+8)":   -948,
		"(1+(4+5+2)-3)+(6+8)":               23,
		" 2-1 + 2 ":                         3,
	}

	for exp, v := range smap {
		calculator := NewCalculator(exp)
		result := calculator.Execuate()
		if result != v {
			t.Errorf("%s result should be %d, but=%d", exp, v, result)
		}
	}

	errmap := map[string]int{
		"(1+(4%51*2)-3)+(6+8)": CalculatorErrorCode["sign_error"],
		"(1+(4/0*2)-3)+(6+8)":  CalculatorErrorCode["divide0"],
		"(1+(45 1*2)-3)+(6+8)": CalculatorErrorCode["empty_error"],
		"1-(     -2))":         CalculatorErrorCode["more)"],
		"1-((     -2)":         CalculatorErrorCode["more("],
	}
	var wg sync.WaitGroup
	for exp, code := range errmap {
		wg.Add(1)
		go func(code int, exp string) {
			calculator := NewCalculator(exp)
			defer func(code int) {
				if r := recover(); r != nil {
					if err, ok := r.(string); ok {
						// 将错误转换为字符串并打印
						if !calculator.CheckPanicErrorCode(err, code) {
							t.Errorf("异常检测失败: %s, %d", err, code)
						}
					} else {
						fmt.Println("捕获到了异常：", r, code)
					}
				}
				wg.Done()
			}(code)
			calculator.Execuate()
		}(code, exp)
	}

	wg.Wait()
}
