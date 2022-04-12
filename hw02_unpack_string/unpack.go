package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Place your code here.
	a := []rune(s)
	var x, b string
	var prev int
	for i := range a {
		letter, err := strconv.Atoi(string(a[i]))
		if err != nil {
			if string(a[i]) == "\\" {
				prev = 2
				continue
			} else {
				if prev == 2 {
					x = x + string(a[i-1]) + string(a[i])
				} else {
					x += string(a[i])
					prev = 1
				}
			}
		} else {
			if prev == 0 {
				return "", ErrInvalidString
			} else {
				if letter == 0 {
					if prev == 2 {
						x = string([]rune(x)[:len(x)-2])
						prev = 0
					} else {
						x = string([]rune(x)[:len(x)-1])
						prev = 0
					}
				} else {
					if prev == 2 {
						b = strings.Repeat((string(a[i-2]) + string(a[i-1])), letter-1)
					} else {
						b = strings.Repeat(string(a[i-1]), letter-1)
					}
					prev = 0
					x += b
				}
			}
		}
	}
	return x, nil
}
