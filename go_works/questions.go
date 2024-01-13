package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Q1 Example
	wordsQ1 := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	resultQ1 := customSort(wordsQ1)
	fmt.Println("Q1:", resultQ1)

	// Q2 Example
	fmt.Print("Q2: ")
	recursiveFunction(9)
	fmt.Println() // Boş bir satır ekleyerek daha iyi görünürlük sağlıyoruz.

	// Q3 Example
	dataQ3 := []string{"apple", "pie", "apple", "red", "red", "red"}
	resultQ3 := mostRepeated(dataQ3)
	fmt.Println("Q3:", resultQ3)
}

// Q1: Sorting by the number of character "a"s and then by length
func customSort(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		// Sort by the number of "a"s in decreasing order
		if countA_i != countA_j {
			return countA_i > countA_j
		}

		// If the same number of "a"s, then sort by length in increasing order
		return len(words[i]) < len(words[j])
	})

	return words
}

// Q2: Recursive function to generate specific output
func recursiveFunction(n int) {
	if n > 0 {
		recursiveFunction(n / 2)
		fmt.Print(n, " ") // Değişiklik burada
	}
}

// Q3: Finding the most repeated data in an array
func mostRepeated(data []string) string {
	countMap := make(map[string]int)
	for _, item := range data {
		countMap[item]++
	}

	var mostRepeatedItem string
	maxCount := 0

	for item, count := range countMap {
		if count > maxCount {
			maxCount = count
			mostRepeatedItem = item
		}
	}

	return mostRepeatedItem
}
