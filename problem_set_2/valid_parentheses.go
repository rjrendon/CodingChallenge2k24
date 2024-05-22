package main
import "fmt"

func isValid(s string) bool {
// Your code here

stack := []rune{}
pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

for _, c := range s {
	if len(stack) > 0 && stack[len(stack)-1] == pairs[c] {
		stack = stack[:len(stack)-1]
	} else {
		stack = append(stack, c)
	}
}

return len(stack) == 0
}


func main() {
// Example usage
inputString := "()[]{}"
result := isValid(inputString)
fmt.Println(result) // Output: true


fmt.Println("--------------------------------------------------------------")
fmt.Println("Test Area")
fmt.Println("--------------------------------------------------------------")
test := "()"
test2 := "()[]{}"
test3 := "(]"
test4 := "([)]"
test5 := "{[]}"
test6 := "{[}" 
fmt.Println(test, "=", isValid(test)) //true
fmt.Println(test2, "=", isValid(test2)) //true
fmt.Println(test3, "=", isValid(test3)) //false
fmt.Println(test4, "=", isValid(test4)) //false
fmt.Println(test5, "=", isValid(test5)) //true
fmt.Println(test6, "=", isValid(test6)) //false

}

