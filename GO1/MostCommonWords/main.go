package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

func GetMostCommonWords(words []string, K int) ([]string, error) {
	if K <= 0 {
		return nil, errors.New("k must be greater than 0")
	}
	if len(words) == 0 {
		return []string{}, nil
	}

	countMap := make(map[string]int)
	for _, word := range words {
		countMap[word]++
	}

	var wordCounts []WordCount
	for word, count := range countMap {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		if wordCounts[i].Count == wordCounts[j].Count {
			return wordCounts[i].Word < wordCounts[j].Word
		} else {
			return wordCounts[i].Count > wordCounts[j].Count
		}
	})

	var result []string
	for i := 0; i < K && i < len(wordCounts); i++ {
		result = append(result, wordCounts[i].Word)
	}

	return result, nil
}

func readInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %v", err)
	}
	return strings.TrimSpace(input), nil
}

func main() {
	input, err := readInput("Input words: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	words := strings.Fields(input)
	var K int
	fmt.Print("Input number: ")
	_, err = fmt.Scanln(&K)
	if err != nil {
		fmt.Printf("Error reading number: %v\n", err)
		return
	}
	result, err := GetMostCommonWords(words, K)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(result) == 0 {
		fmt.Println("No words found.")
	} else {
		fmt.Println(strings.Join(result, " "))
	}
}
