package goes

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

type camelSnakeStateMachine int

const (
	idle camelSnakeStateMachine = iota
	firstAlphaNum
	alphaNum
	delimiter
)

func (s camelSnakeStateMachine) next(r rune) camelSnakeStateMachine {
	switch s {
	case idle:
		if isAlphaNum(r) {
			return firstAlphaNum
		}
	case firstAlphaNum:
		if isAlphaNum(r) {
			return alphaNum
		}
		return delimiter
	case alphaNum:
		if !isAlphaNum(r) {
			return delimiter
		}
	case delimiter:
		if isAlphaNum(r) {
			return firstAlphaNum
		}
		return idle
	}
	return s
}

func CamelCase(str string) string {
	var b strings.Builder

	stateMachine := idle
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		i += size
		stateMachine = stateMachine.next(r)
		switch stateMachine {
		case firstAlphaNum:
			b.WriteRune(unicode.ToUpper(r))
		case alphaNum:
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func SnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func InitialLowerCase(str string) string {
	a := []rune(str)
	a[0] = unicode.ToLower(a[0])
	str = string(a)

	return str
}

func isAlphaNum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}
