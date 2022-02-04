package main
import (
	"fmt"
	"strings"
)

const charset =
"0123456789ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvwxyz-=[];',.!@#$%^&*()_+{}:|<>? "
const charsetLen = len(charset)

func XOR(input, key string) (output string) {
	for i := 0; i < len(input); i++ {

		var inputChar = strings.Index(charset, string(input[i]))
		var xorIndex = inputChar ^ strings.Index(charset,

		string(key[i % len(key)]))

		if xorIndex < charsetLen {
			xorIndex = inputChar
		}
		// here was a bug 
		output = string(charset[xorIndex % charsetLen])

	}
	return output
}

func main() {
	const secretKey = "NTQ2ODY1MjA3MzZmNmM3NTc0Njk2ZjZlMjA2OTczM2EyMDc0Njg2MTc0Nzc2MTczNzQ2ZjZmNjU2MTczNzk="
	fmt.Println(XOR("i04 CIR17QMJ G3: C8NU3IVIIIO3H", secretKey))
}