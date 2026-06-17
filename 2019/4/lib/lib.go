package lib

import (
	"strconv"
)

func IsPossiblePassword(password int) (bool, bool) {
	s := strconv.Itoa(password)
	hasDouble := false
	hasOnlyDouble := false
	ascending := true
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			hasDouble = true
			j := i + 1
			if j >= len(s) || s[j] != s[i] {
				if i <= 1 || s[i-2] != s[i] {
					hasOnlyDouble = true
				}
			}
		}
		if s[i] < s[i-1] {
			ascending = false
		}
	}
	return hasDouble && ascending, hasOnlyDouble && ascending
}
