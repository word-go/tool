package tool

import "regexp"

/**
字符串截取
*/
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

/**
验证手机号格式
*/
func CheckMobile(mobile string) bool {
	regular := "/^13\\d{9}$|^14\\d{9}$|^15\\d{9}$|^17\\d{9}$|^18\\d{9}$/"
	return regexp.MustCompile(regular).MatchString(mobile)
}
