package dca

import "fmt"

// first test of package and github sync
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
