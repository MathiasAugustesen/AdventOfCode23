package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v2"
)

func part1() {
	inputFile, _ := os.ReadFile("input.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(inputFile)))
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		barIndex := strings.Index(line, "|")
		winningNumbersSlice := line[colonIndex + 2:barIndex - 1]
		winningNumbers := set.New[int64](16)
		for _, numStr := range strings.Fields(winningNumbersSlice) {
			parsedNum, _ := strconv.ParseInt(numStr, 10, 64)
			winningNumbers.Insert(parsedNum)
		}

		cardSlice := line[barIndex + 2:]
		cardNumbers := set.New[int64](16)
		for _, numStr := range strings.Fields(cardSlice) {
			parsedNum, _ := strconv.ParseInt(numStr, 10, 64)
			cardNumbers.Insert(parsedNum)
		}
		
		intersectionCount := cardNumbers.Intersect(winningNumbers).Size()
		sum += calculateScore(intersectionCount)
	}
	fmt.Println(sum)

}

func calculateScore(count int) int {
	if count == 0 {
		return 0
	}
	return 1 << (count - 1)
}

func part2() {
	inputFile, _ := os.ReadFile("input.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(inputFile)))
	scores := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		barIndex := strings.Index(line, "|")

		winningNumbersSlice := line[colonIndex + 2:barIndex - 1]
		winningNumbers := set.New[int64](16)
		
		for _, numStr := range strings.Fields(winningNumbersSlice) {
			parsedNum, _ := strconv.ParseInt(numStr, 10, 64)
			winningNumbers.Insert(parsedNum)
		}

		cardSlice := line[barIndex + 2:]
		cardNumbers := set.New[int64](16)
		for _, numStr := range strings.Fields(cardSlice) {
			parsedNum, _ := strconv.ParseInt(numStr, 10, 64)
			cardNumbers.Insert(parsedNum)
		}
		
		intersectionCount := cardNumbers.Intersect(winningNumbers).Size()
		scores = append(scores, intersectionCount)
	}

	cardCounts := make([]int, len(scores))
	for i := 0; i < len(scores); i++ {
		cardCounts[i] = 1
	}

	for i, score := range scores {
		currentCards := cardCounts[i]
		for j := i + 1; j < int(math.Min(float64(i + score + 1), float64(len(scores)))); j++ {
			cardCounts[j] += currentCards
		}
	}
	sum := 0
	for _, cards := range cardCounts {
		sum += cards
	}
	fmt.Println(sum)
}