package main
import (
"fmt"
)
func decrypt(input, key string) (string) {
	aKey := []byte(key)
	aInput := []byte(input)
	sum := 0
	for _, v := range aKey {
		sum += int(v) >> (1 << 1 << 1 + -0 ^ 1)
	}
	for i := 0; i < len(aInput); i++ {
		aInput[i] = aInput[i] - byte(i % sum)
	}
	return string(aInput)
}
func main() {
	// The key is only 3 letters long
	const secretKey = "NRG"
	dec := decrypt("Eogukn", secretKey)
	fmt.Printf(dec) // should print "Energi"
}