package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type DisplacementParams struct {
	acceleration        float64
	initialVelocity     float64
	initialDisplacement float64
}

func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	f := func(t float64) float64 {
		return a*(math.Pow(t, 2)*0.5) + (v * t) + s
	}
	return f
}

func validateFloatOrExit(input string) float64 {
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error:", err, "Exiting")
		os.Exit(1)
	}
	return result
}

func getUserInput() string {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	return input
}

func getDisplacementParams() DisplacementParams {
	var input string
	var acc float64
	fmt.Println("Enter acceleration:")
	input = getUserInput()
	acc = validateFloatOrExit(input)

	var initVel float64
	fmt.Println("Enter initial velocity:")
	input = getUserInput()
	initVel = validateFloatOrExit(input)

	var initDispl float64
	fmt.Println("Enter initial displacement:")
	input = getUserInput()
	initDispl = validateFloatOrExit(input)

	return DisplacementParams{acceleration: acc, initialVelocity: initVel, initialDisplacement: initDispl}
}

func main() {
	result := getDisplacementParams()
	fmt.Println(result)

	//get the displacement function as a function of time
	fn := GenDisplaceFn(result.acceleration, result.initialVelocity, result.initialDisplacement)

	for {
		fmt.Println("Please enter a value for time (x to quit): ")
		input := getUserInput()
		if strings.ToLower(input) == "x" {
			break
		} else {
			t := validateFloatOrExit(input)

			//call the displacement function with each value of time
			displacement := fn(t)

			fmt.Printf("Displace at time %f is %f\n", t, displacement)
		}
	}
}
