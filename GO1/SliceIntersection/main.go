package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SliceIntersection(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	countMap := make(map[int]int)
	for _, num := range nums2 {
		countMap[num]++
	}

	result := []int{}
	for _, num := range nums1 {
		if count, ok := countMap[num]; ok && count > 0 {
			result = append(result, num)
			countMap[num]--
		}
	}

	return result
}

func Reader(nums *[]int) error {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return fmt.Errorf("Invalid input")
			}
			*nums = append(*nums, num)
		}
	}
	return nil
}

func main() {
	var nums1, nums2 []int
	fmt.Println("Enter the first numbers array:")
	err := Reader(&nums1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter the second numbers array:")
	err = Reader(&nums2)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := SliceIntersection(nums1, nums2)
	if len(result) == 0 {
		fmt.Println("Empty intersection")
	} else {
		fmt.Println("Intersection slice:", result)
	}
}
