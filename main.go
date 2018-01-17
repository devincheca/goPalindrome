package main
import
(
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	ch1 := make(chan byte)
	ch2 := make(chan byte)
	isPalindrome := true
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter palindrome to check:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.Replace(input, "\n", "", -1)
	go ReadForward(ch1, input)
	go ReadBackward(ch2, input)
	for i := 0; i < len(input); i++ {
		if <-ch1 != <-ch2 {
			isPalindrome = false
		}
	}
	if isPalindrome {
		fmt.Println(input, "is a palindrome.")
	} else {
		fmt.Println(input, "is not a palindrome.")
	}
}
func ReadForward(ch chan byte, input string) {
	for i := 0; i < len(input); i++ {
		ch <- input[i]
	}
}
func ReadBackward(ch chan byte, input string) {
	for i := len(input); i > 0; i-- {
		ch <- input[i - 1]
	}
}
