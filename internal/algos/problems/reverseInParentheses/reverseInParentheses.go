package reverseinparentheses

import "fmt"

func ReverseInParentheses(inputString string) string {
	fmt.Println("Input: ", inputString)
	res := []byte{}

	for i := 0; i < len(inputString); i++ {
		fmt.Printf("\ni: %v; Value at i: %s; Current res: %s\n", i, string(inputString[i]), res)
		if string(inputString[i]) == ")" {
			popped, result := popTillOpeningBracket(res)
			fmt.Printf("Popped: %s; Current res: %s\n", popped, res)
			result = append(result, popped...)
			res = result
			fmt.Printf("Res append after popped: %s\n", res)
		} else {
			res = append(res, inputString[i])
		}
	}
	return string(res)
}

func popTillOpeningBracket(input []byte) ([]byte, []byte) {
	fmt.Printf("popTillOpeningBracket input: %s\n", input)
	res := []byte{}
	for i := len(input) - 1; i >= 0; i-- {
		if string(input[i]) == "(" {
			input = input[:i]
			break
		}
		res = append(res, input[i])
	}
	return res, input
}
