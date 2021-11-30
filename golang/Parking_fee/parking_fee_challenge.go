// The challenge is to write code without
// using if or else (author Rik Wittkopp).

package main

import "fmt"

type parking_fees []struct {
	minutes, rate int
}

func (pf parking_fees) CalculateFeeCost(parking_minutes int) int {

	// First element of the array will
	// store the parking fee rate
	// in case of a false comparison.
	// Second element of the array will
	// store the parking fee rate
	// in case of a true comparison.
	var fee_rates [2]int

	// There are no conversions
	// from bool to integer types
	// (https://stackoverflow.com/questions/8393933/is-there-a-way-to-convert-integers-to-bools-in-go-or-vice-versa)
	// Map of boolean and integer
	// using built-in associative data type.
	var bool2int = map[bool]int{false: 0, true: 1}

	// Compare parking minutes with
	// the minutes from parking fee struct
	for _, parking_fee := range pf {
		// Assign parking fee rate
		// to fee_rates array
		fee_rates[bool2int[parking_minutes >= parking_fee.minutes]] = parking_fee.rate
	}

	// Calculate fee cost
	// Value of the calculated parking fee
	// rate is in the second element of the
	// fee_rates array.
	rate := fee_rates[1]
	coef := parking_minutes / 60
	hours := coef + bool2int[parking_minutes > 60*coef]
	return rate * hours

}

func main() {
	// There are four categories of fees:
	// 1. first 20 minutes is free
	// 2. from 21 to 60 minutes is 10$
	// 3. from 61 to 179 minutes is 20$
	// per hour or part thereof.
	// 4. from 180 minutes and more is 40$
	// per hour or part thereof.
	pf := parking_fees{{minutes: 0, rate: 0}, {21, 10}, {61, 20}, {180, 40}}

	for parking_minutes := 0; parking_minutes <= 201; parking_minutes += 3 {
		fmt.Printf("Parking minutes %v -> %v$\n", parking_minutes, pf.CalculateFeeCost(parking_minutes))
	}
}


