package lib

func Intcode(input []int) []int {
	operator := input[0]
	a := input[1]
	b := input[2]
	position := input[3]
	if operator == 1 {
		input[position] = input[a] + input[b]
	}
	if operator == 2 {
		input[position] = input[a] * input[b]
	}
	return input
}
