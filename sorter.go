package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func itemSort(items []int, c chan []int) {
	sort.Ints(items)
	c <- items
}

func itemStrToInt(itemStr []string) ([]int, error) {
	var itemInts []int
	for _, item := range itemStr {
		result, err := strconv.Atoi(item)
		if err != nil {
			return itemInts, err
		}
		itemInts = append(itemInts, result)
	}
	return itemInts, nil
}

func splitItemsToGroups(items []int, numOfGroups int) [][]int {
	var results [][]int
	if len(items) < numOfGroups {
		log.Fatal("Cannot have less items than the number of groups")
	}
	for i := 0; i < numOfGroups; i++ {
		lower := i * len(items) / numOfGroups
		upper := ((i + 1) * len(items)) / numOfGroups
		results = append(results, items[lower:upper])
	}
	return results
}

func main() {
	fmt.Println("Enter some integers separated by a space (at least 4 integers):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal("Error with input")
	}

	itemsString := strings.Split(strings.TrimSpace(scanner.Text()), " ")

	results, err := itemStrToInt(itemsString) // get the slice of integers
	if err != nil {
		log.Fatal("Error converting input, invalid integers")
	}

	groups := splitItemsToGroups(results, 4) // create 4 groups of items
	sortedResultsChan := make(chan []int)

	for _, g := range groups {
		go itemSort(g, sortedResultsChan) // sort each group
	}

	var aggregatedResults []int

	for i := 0; i < 4; i++ {
		items := <-sortedResultsChan
		fmt.Printf("Result received: %v\n", items)
		aggregatedResults = append(aggregatedResults, items...)
	}

	fmt.Println("Performing final sort")
	go itemSort(aggregatedResults, sortedResultsChan) // sort final slice
	fmt.Printf("Result received: %v\n", <-sortedResultsChan)
}
