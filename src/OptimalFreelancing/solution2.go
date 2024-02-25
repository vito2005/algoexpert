package main

import (
	"fmt"
)

// time O(n) | space O(1)
func OptimalFreelancing(jobs []map[string]int) int {
	var profit int
	MAX_DAYS_LENGTH := 7
	var completedJobs = []map[string]int{}

	for day := MAX_DAYS_LENGTH; day > 0; day-- {
		var bestJob map[string]int
		for _, job := range jobs {
			if !has(job, completedJobs) &&
				job["deadline"] >= day &&
				bestJob["payment"] <= job["payment"] {
				bestJob = job
			}
		}
		if bestJob != nil {
			completedJobs = append(completedJobs, bestJob)
		}
		profit += bestJob["payment"]
	}

	return profit
}

func main() {

	input := []map[string]int{
		{
			"deadline": 2,
			"payment":  2,
		},
		{
			"deadline": 4,
			"payment":  3,
		},
		{
			"deadline": 5,
			"payment":  1,
		},
		{
			"deadline": 7,
			"payment":  2,
		},
		{
			"deadline": 3,
			"payment":  1,
		},
		{
			"deadline": 3,
			"payment":  2,
		},
		{
			"deadline": 1,
			"payment":  3,
		},
	}

	expected := 13
	actual := OptimalFreelancing(input)

	fmt.Println(actual == expected)
}

func has(item map[string]int, arr []map[string]int) bool {
	for _, val := range arr {
		if val["payment"] == item["payment"] && val["deadline"] == item["deadline"] {
			return true
		}
	}
	return false
}
