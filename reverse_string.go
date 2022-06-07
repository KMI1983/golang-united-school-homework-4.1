package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	result := []rune("")
	for i := len(runes) - 1; i >= 0; i-- {
		result = append(result, runes[i])
	}
	return string(result)
}
