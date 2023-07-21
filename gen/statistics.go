/*
 * File: statistics.go
 * Description: displays statistics about the build proccess.
 */


package gen

import "fmt"

// keeps track of count
var successes int
var skipped int

// increments the successes variable when a function is successfully wrapped
func IncrementSuccess() {
	successes++
}

// increments the skipped variable when a function is skipped
func IncrementSkipped() {
	skipped++
}

// prints statistics about the build process
func PrintStats() {
	fmt.Println("\n\nStatistics:")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Successfully wrapped methods: %8d", successes)
	fmt.Printf("\nSkipped methods: %21d\n\n", skipped)
}
