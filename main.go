/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2025-11-12
 * This file finds and displays the area of a circle with a radius of 5.6 cm.
 */

package main

import "fmt"

func main() {
	// int data type radius
	var radius float32 = 5.6

	// stores 3.14159 into PI
	const PI float32 = 3.14159

	// INPUT - none

	// PROCESS
	// calculate area
	answer := (PI * (radius * radius))

	// OUTPUT
	// display the result
	fmt.Println("The area of a circle with a radius of 5.6 cm is", answer, "cm².")

	fmt.Println("\nDone.")
}
