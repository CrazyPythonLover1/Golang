package main
import (
"fmt"
"math/rand"
"time"
)
func isPalindrome(number int) (bool) {
	if number < 0 {
		return false
	}

	tmp := number
	var valueList []int
	times := 1

	for tmp != 0 {
		valueList = append(valueList, tmp%10)
		tmp /= 10
	}

	tmp = 0
	for i := len(valueList) - 1; i >= 0; i-- {
		tmp += valueList[i] * times
		times *= 10
	}

	return number == tmp
}

func main(){
	rand.Seed(time.Now().UnixNano())
	var number int = rand.Intn(10000)
	if isPalindrome(number) {
		fmt.Printf("%d is a palindrome number", number)
		fmt.Println()
	} else {
		fmt.Printf("%d is NOT a palindrome number", number)
		fmt.Println()
	}
}