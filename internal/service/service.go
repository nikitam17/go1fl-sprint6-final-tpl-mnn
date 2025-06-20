package service

import (
	"go1fl-sprint6-final-tpl-mnn/pkg/morse"
	"strings"
)

func Conver(s string) string {
	strData := string(s)
	if isMorseCode(strData) {
		return morse.ToText(strData)
	}
	return morse.ToMorse(strData)
}

func isMorseCode(s string) bool {
	return strings.ContainsFunc(s, isMorseChar)
}

func isMorseChar(r rune) bool {
	if r == '.' || r == '-' {
		return true
	}
	return false // Если найден символ, отличный от точек и тире, это не Морзе
}
