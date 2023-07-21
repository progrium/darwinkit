/*
 * File: statistics.go
 * Description: displays statistics about the build proccess.
 */


package gen

import "fmt"

// keeps track of count
var successes int
var skipped int
var classes int

// increments the successes variable when a function is successfully wrapped
func IncrementSuccess() {
	successes++
}

// increments the skipped variable when a function is skipped
func IncrementSkipped() {
	skipped++
}

func IncrementClasses() {
	classes++
}

// prints statistics about the build process
func PrintStats() {
	fmt.Println("\n\nStatistics:")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("%-33s %d", "Wrapped methods:", successes)
	fmt.Printf("\n%-33s %d", "Skipped methods:", skipped)
	fmt.Printf("\n%-33s %.2f%%", "Percent wrapped:", float64(successes)/float64(successes + skipped)*100)
	fmt.Printf("\n%-33s %.2f%%", "Percent skipped:", float64(skipped)/float64(successes + skipped)*100)
	fmt.Printf("\n%-33s %d\n\n", "Classes:", classes)
}
