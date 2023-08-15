package composeRanges

import (
	"strconv"
)

type Range struct {
	Start string
	End   string
}

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

func Compose(nums []int, results []Range) []Range {

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
