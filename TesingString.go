package main

import (
	"fmt"
	"regexp"
)

func main() {
	string := "Hello"
	re := regexp.MustCompile(`\b` + string + `\b`)

	matched := re.MatchString("Hello World")
	fmt.Println("Matched Found =", matched)
}
