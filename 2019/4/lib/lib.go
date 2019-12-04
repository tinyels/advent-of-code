package lib

import (
	"strconv"
)

func IsPossiblePassword(password int) bool {
	s := strconv.Itoa(password)
	hasDouble := false
	ascending := true
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			hasDouble = true
		}
		if s[i] < s[i-1] {
			ascending = false
		}
	}
	return hasDouble && ascending
}
