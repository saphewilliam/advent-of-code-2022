package lib

import "strconv"

type Solution struct {
	I int
	S string
}

func (s *Solution) GetValue() string {
	if s.I != 0 {
		return strconv.Itoa(s.I)
	}
	return s.S
}

// Create an integer-type score
func IScore(score int) Solution {
	return Solution{I: score}
}

// Create a string-type score
func SScore(score string) Solution {
	return Solution{S: score}
}
