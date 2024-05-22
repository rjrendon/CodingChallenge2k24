package main
import "fmt"

//Palindrome checker
func isPalindrome(s string) bool {
// Your code here

for i := 0; i < len(s)/2; i++ {
	if s[i] != s[len(s)-1-i] {
		return false
	}
}
return true

}


func palindromePairs(words []string) [][]int {
// Your code here
var result [][]int
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && isPalindrome(words[i]+words[j]) {
				result = append(result, []int{i, j})
			}
		}
	}
	return result

}
func main() {
// Example usage
inputWords := []string{"bat", "tab", "cat"}
result := palindromePairs(inputWords)
fmt.Println(result) // Output: [[0 1] [1 0]]


fmt.Println("--------------------------------------------------------------")
fmt.Println("Test Area")
fmt.Println("--------------------------------------------------------------")

//palindrome function test
str := "radar"
str2 := "racedriver"
res := isPalindrome(str)
res2 := isPalindrome(str2)
fmt.Println(str, "=", res) // true
fmt.Println(str, "=", res2) //false

//different inputs to ensure it handles various cases.
inputWords2 := []string{"rat", "bar", "tar"}
result2 := palindromePairs(inputWords2)
fmt.Println(inputWords2, "=", result2) 

inputWords3 := []string{"lor", "nor", "ron"}
result3 := palindromePairs(inputWords3)
fmt.Println(inputWords3, "=", result3) 

}

