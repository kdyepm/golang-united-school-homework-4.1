package reverse_string

func ReverseString(input string) (output string) {
	rns := []rune(input)
	for i := 0; i < (len(rns))/2; i++ {
		rns[i], rns[len(rns)-i-1] = rns[len(rns)-i-1], rns[i]
	}
	output = string(rns)
	return output
}
