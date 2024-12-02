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

func parseIntField(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return num
}

func parseLists() ([]int, []int) {
	var a, b []int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 2 {
			a = append(a, parseIntField(parts[0]))
			b = append(b, parseIntField(parts[1]))
		}
	}
	return a, b
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func measureDistance(a, b []int) {
	distance := 0
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	for i := range a {
		distance += abs(a[i] - b[i])
	}

	fmt.Printf("distance: %d\n", distance)
}

func measureSimilarityScore(a, b []int) {
	score := 0
	for _, k := range a {
		for _, j := range b {
			if k == j {
				score += k
			}
		}
	}
	fmt.Printf("similarity score: %d\n", score)
}

func main() {
	first, second := parseLists()
	measureDistance(first, second)
	measureSimilarityScore(first, second)
}
