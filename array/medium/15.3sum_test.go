package medium

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/3sum/description/
// * the description show that [-1,0,1], [1,0,-1], and [0,1,-1] are same, due to the order does not matter

// method 1 use two for loop
// 1) sort the nums, TC = O(NlogN), SC = O(logN)
// 2) use two nested for loop to scan the nums, and use result to store the three sum values
// 3) during the second for loop process, we need to store every complement into the numMap (this numMap will be refresh, when turn into the next loop of in first for loop)
// TC = O(N^2), SC = O(N)
func threeSum1(nums []int) [][]int {
	// * this sort step is essential, if we want to eliminate the repeat from [-1,0,1], [1,0,-1], and [0,1,-1] during the following steps. Finally only has one result like [-1,0,1]
	sort.Ints(nums) // TC = O(NlogN), SC = O(logN)

	result := [][]int{} // SC = O(N)

	for i := 0; i < len(nums)-2; i++ {
		// this step check the current value with previous value, if there are same, mean the same result already exist. Then skip the repeat.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		numMap := map[int]bool{}

		for j := i + 1; j < len(nums); j++ {
			complement := -(nums[i] + nums[j])

			if numMap[complement] {
				result = append(result, []int{nums[i], complement, nums[j]})

				// if nums[j] == nums[j+1] == 1; mean result of nums[j]  is [-1,0,1]; result of nums[j+1] is same as [-1,0,1]
				// this step check the "current value" with "next value", mean the same result already exist. Then skip the repeat.
				for j < len(nums)-1 && nums[j] == nums[j+1] { // len(nums)-1 is the last second number; last number len(nums) don't have nums[j+1]
					j++
				}
			}

			numMap[nums[j]] = true
		}
	}

	return result
}

// method 2 use two nested for loop and two pointer
// the steps reason are almost same as method 1, but thie method 2 use two pointer to replace the numMap, and the two pointer will be refresh when turn into the next loop of in first for loop
// TC = O(N^2), SC = O(N)
// although the complexity are same as method 1, but this method 2 get the best performance than method 1 in leetcode !!!
// * this is the best solution for me currently
func threeSum2(nums []int) [][]int {
	// * this sort step is essential, if we want to eliminate the repeat from [-1,0,1], [1,0,-1], and [0,1,-1] during the following steps. Finally only has one result like [-1,0,1]
	sort.Ints(nums) // TC = O(NlogN), SC = O(logN)

	result := [][]int{} // SC = O(N)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		leftIdx, rightIdx := i+1, len(nums)-1

		for leftIdx < rightIdx {
			sum := nums[i] + nums[leftIdx] + nums[rightIdx]

			if sum < 0 {
				// due to the nums is sorted, so if sum < 0, mean the nums[leftIdx] is too small, need to increase the nums[leftIdx]
				leftIdx++
			} else if sum > 0 {
				// due to the nums is sorted, so if sum > 0, mean the nums[rightIdx] is too big, need to decrease the nums[rightIdx]
				rightIdx--
			} else {
				result = append(result, []int{nums[i], nums[leftIdx], nums[rightIdx]})

				// if nums[leftIdx] == nums[leftIdx+1] == 1; mean result of nums[leftIdx]  is [-1,0,1]; result of nums[leftIdx+1] is same as [-1,0,1]
				// this step check the "current value" with "next value", mean the same result already exist. Then skip the repeat.
				for leftIdx < rightIdx && nums[leftIdx] == nums[leftIdx+1] {
					leftIdx++
				}

				// if nums[rightIdx] == nums[rightIdx-1] == 1; mean result of nums[rightIdx]  is [-1,0,1]; result of nums[rightIdx-1] is same as [-1,0,1]
				// this step check the "current value" with "next value", mean the same result already exist. Then skip the repeat.
				for leftIdx < rightIdx && nums[rightIdx] == nums[rightIdx-1] {
					rightIdx--
				}

				leftIdx++
				rightIdx--
			}
		}
	}

	return result
}

func Test_threeSum1(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result [][]int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			expected: expected{
				result: [][]int{{-1, 0, 1}, {-1, -1, 2}},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 1, 1},
			},
			expected: expected{
				result: [][]int{},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 0, 0},
			},
			expected: expected{
				result: [][]int{{0, 0, 0}},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			threeSum1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_threeSum2(t *testing.T) {
	type args struct {
		nums []int
	}
	type expected struct {
		result [][]int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			expected: expected{
				result: [][]int{{-1, -1, 2}, {-1, 0, 1}},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 1, 1},
			},
			expected: expected{
				result: [][]int{},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{0, 0, 0},
			},
			expected: expected{
				result: [][]int{{0, 0, 0}},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			threeSum2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
