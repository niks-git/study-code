package main

import "fmt"

func main() {
	fmt.Println("LeetCode: Problem 1")
	fmt.Println("Two Sum")
}

func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, v := range nums { // O(n) for this loop
		complement := target - v
		index, ok := hash[complement]
		if ok {
			return []int{index, i}
		}
		hash[v] = i
	}
	return nil
}
