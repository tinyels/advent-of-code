package lib

func Intcode(input []int) []int {
	var i, a, b, outputPosition int
	for input[i] != 99 {
		a = input[i+1]
		b = input[i+2]
		outputPosition = input[i+3]
		switch input[i] {
		case 1:
			input[outputPosition] = input[a] + input[b]
		case 2:
			input[outputPosition] = input[a] * input[b]
		}
		i += 4
	}
	return input
}
