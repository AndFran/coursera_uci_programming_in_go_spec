package main

import (
	"fmt"
	"os"
	"strconv"
)

func Swap(integers []int, i int) {
	integers[i], integers[i+1] = integers[i+1], integers[i]
}

func BubbleSort(toSort []int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(toSort)-1; i++ {
			if toSort[i] > toSort[i+1] {
				Swap(toSort, i)
				swap = true
			}
		}
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

func getUserInput() []int {
	var input string
	var number int
	results := make([]int, 0, 10)

	fmt.Println("Please enter 10 numbers")
	for i := 0; i < 10; i++ {
		fmt.Printf("Please enter number %d of 10:", i+1)

		_, err := fmt.Scan(&input)
		exitOnError(err)

		number, err = strconv.Atoi(input)
		exitOnError(err)

		results = append(results, number)
	}
	return results
}

func main() {
	results := getUserInput()
	fmt.Println("Starting results set:", results)

	BubbleSort(results)

	fmt.Println("Sorted results set:", results)

}
