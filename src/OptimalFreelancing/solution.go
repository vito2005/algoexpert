package main

import (
	"fmt"
	"sort"
)

// time O(nlogn) | space O(1)
func OptimalFreelancing(jobs []map[string]int) int {
	var profit int
	MAX_DAYS_LENGTH := 7

	timeline := make([]bool, MAX_DAYS_LENGTH)

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i]["payment"] >= jobs[j]["payment"]
	})
	fmt.Println(jobs)

	for _, job := range jobs {
		for time := min(job["deadline"], MAX_DAYS_LENGTH) - 1; time >= 0; time-- {
			if timeline[time] == false {
				timeline[time] = true
				profit += job["payment"]
				break
			}
		}
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
