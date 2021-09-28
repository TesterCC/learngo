package word

import (
	"strings"
	"unicode"
)

// P24 1.单词全部转为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 1.单词全部转为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// P25 2.下画线单词转大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 3.下画线单词转小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// P25 use govalidator   4.驼峰单词转下划线单词

func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
