package main

import "fmt"

// Adapted rom https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
// and Russ Cox, on the golang-nuts mailing list (https://groups.google.com/forum/#!topic/golang-nuts/oPuBaYJ17t4)
func main() {
	input := "The quick brown 狐 jumped over the lazy 犬"
	fmt.Println("input : " + input)
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	fmt.Println("output : " + output)
}
