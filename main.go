/*
 * @author Bianca Bertinato
 * @version 1.0.0
 * @date 2025-10-08
 * @fileoverview This program displays the average of three numbers.
 */
package main

import "fmt"

func main() {

    // initialize numbers as constant
		const NUMBER1 float64 = 56.9
		const NUMBER2 float64 = 89.7
		const NUMBER3 float64 = 90.2

		// PROCESS
		// calculate the average
    answer := (NUMBER1 + NUMBER2 + NUMBER3) / 3

		// OUTPUT
		// display the result
		fmt.Println("The average of", NUMBER1,",", NUMBER2, "and", NUMBER3, "is", answer)

		fmt.Println("\nDone.")
}		
