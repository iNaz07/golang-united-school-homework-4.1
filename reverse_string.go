package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		output += string(runes[i])
	}
	return output
}

// func main() {
// 	fmt.Println(ReverseString("Hello world!"))
// 	fmt.Println(ReverseString("Hello\n world!"))
// 	fmt.Println(ReverseString("Señor"))
// }
