package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countOccurrnces(leftList, rightList []int) map[int]int {
	occurences := map[int]int{}

	for _, val := range leftList {
		for _, rightVal := range rightList {
			if val == rightVal {
				occurences[val]++
			}
		}
	}

	return occurences
}

func reduce(arr []int) int {
	result := 0
	for _, val := range arr {
		result += val
	}
	return result
}

func main() {
	const input = "./input.txt"
	content, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	leftList, rightList := make([]int, 0), make([]int, 0)
	for _, val := range lines {
		data := strings.Split(val, "   ")

		firstIntVal, _ := strconv.Atoi(data[0])
		secondIntVal, _ := strconv.Atoi(data[1])

		leftList = append(leftList, firstIntVal)
		rightList = append(rightList, secondIntVal)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	allDistances := make([]int, 0)
	for index, val := range leftList {
		rightListVal := rightList[index]

		allDistances = append(allDistances, int(math.Abs(float64(val-rightListVal))))
	}

	log.Printf("The sum of the differences is: %d", reduce(allDistances))

	occurrences := countOccurrnces(leftList, rightList)

	similarityScores := make([]int, 0)
	for k, v := range occurrences {
		similarityScores = append(similarityScores, k*v)
	}

	log.Printf("The similarity score is: %d", reduce(similarityScores))
}
