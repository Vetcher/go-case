package go_case

import (
	"strings"
	"unicode"
)

var (
	// ToSnakeCase converts string to snake_case.
	ToSnakeCase = ToSomeCaseWithSep('_', unicode.ToLower)
	// ToLowerCase is a shortcut for strings.ToLower.
	ToLowerCase = strings.ToLower
	// ToUrlSnakeCase works like ToSnakeCase except fills spaces with `-`.
	ToUrlSnakeCase = ToSomeCaseWithSep('-', unicode.ToLower)
	// ToDotSnakeCase works like ToSnakeCase except fills spaces with `.`(dot).
	ToDotSnakeCase = ToSomeCaseWithSep('.', unicode.ToLower)
	// ToCamelCase converts string to CamelCase. All space symbols, `_`, `.` and `-` will be removed. First symbol will be capital.
	ToCamelCase = toCamelCase(true, unicode.ToUpper)
	// ToCamelCase converts string to CamelCase. All space symbols, `_`, `.` and `-` will be removed. First symbol will be capital.
	ToCamelCaseLowerFirst = toCamelCase(false, unicode.ToUpper)
	// ToNoCase returns input's string. Should be used for mocking, tests, etc.
	ToNoCase = func(s string) string { return s }
)

// ToSomeCaseWithSep allow to get customized converter, that will work as ToSnakeCase, but with your rules.
func ToSomeCaseWithSep(sep rune, runeConv func(rune) rune) func(string) string {
	return func(s string) string {
		in := []rune(s)
		n := len(in)
		var runes []rune
		for i, r := range in {
			if isExtendedSpace(r) {
				runes = append(runes, sep)
				continue
			}
			if unicode.IsUpper(r) {
				if i > 0 && sep != runes[i-1] && ((i+1 < n && unicode.IsLower(in[i+1])) || unicode.IsLower(in[i-1])) {
					runes = append(runes, sep)
				}
				r = runeConv(r)
			}
			runes = append(runes, r)
		}
		return string(runes)
	}
}

func toCamelCase(upperFirst bool, runeConv func(rune) rune) func(string) string {
	return func(s string) string {
		in := []rune(s)
		var runes []rune
		for i, r := range in {
			if isExtendedSpace(r) {
				continue
			}
			if unicode.IsUpper(r) {
				runes = append(runes, r)
				continue
			}
			if i > 0 && isExtendedSpace(in[i-1]) && unicode.IsLower(r) || i == 0 && upperFirst {
				r = runeConv(r)
			}
			runes = append(runes, r)
		}
		return string(runes)
	}
}

func isExtendedSpace(r rune) bool {
	return unicode.IsSpace(r) || r == '_' || r == '-' || r == '.'
}
