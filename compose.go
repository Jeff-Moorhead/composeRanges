package composeRanges

/*
Problem statement:

Given a sorted integer array that does not contain any duplicates, return a summary of the number ranges it contains.

Example:

For nums = [-1, 0, 1, 2, 6, 7, 9], the output should be
	solution(nums) = ["-1->2", "6->7", "9"].
*/

import (
	"strconv"
)

// A convenient wrapper for ranges (much cleaner than a map or array)
type Range struct {
	Start string
	End   string
}

// This function just glues the final results together
func Solution(nums []int) []string {

	ranges := Compose(nums, nil)
	results := make([]string, 0)

	for _, r := range ranges {
		if r.End == "" {
			results = append(results, r.Start)
		} else {
			final := r.Start + "->" + r.End
			results = append(results, final)
		}
	}

	return results
}

// This is the actual recursive solution to the DP problem
func Compose(nums []int, results []Range) []Range {

	// TODO: how can we use memoization here? Is it necessary?
	if results == nil {
		results = make([]Range, 0)
	}

	if len(nums) == 0 {
		return results
	}

	start := 0
	end := len(nums) - 1

	for end > start {

		// For consecutive numbers in an array, the last element minus the first element SHOULD equal the final array index (unproven)
		// TODO: is there a way to prove rigorously that the above statement is true for all consecutive arrays?
		if nums[end]-nums[start] == end {
			final := Range{
				Start: strconv.Itoa(nums[start]),
				End:   strconv.Itoa(nums[end]),
			}

			results = append(results, final)

			return Compose(nums[end+1:], results)
		} else {
			end--
			continue
		}
	}

	final := Range{
		Start: strconv.Itoa(nums[start]),
	}

	results = append(results, final)

	return Compose(nums[start+1:], results)
}
