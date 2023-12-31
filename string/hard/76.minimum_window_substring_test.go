package hard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-window-substring/description/

// method 1 sliding window + hash table
// 1) use a tCharCountMap to store the characters of t (key: character, value: count)
// 2) use a wCharCountMap to store the characters of s (key: character, value: count) in the current window
// 3) initial sliding window, use two pointers, left and right, left = 0, right = 0, for each iteration, add the character of the right of the window to the wCharCountMap
// 4) count == 0 means the current window contains all the characters of t, then move the left pointer to shrink the window, and update the wCharCountMap, update the minLength, update the start and end index of the current window for result
// 5) check minLength, if minLength no updated, it means there is no window contains all the characters of t
// 6) finally, return the result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func minWindow1(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// use a map to store the characters of t
	/*
		s = "ADOBECODEBANC"
		t = "ABC"
		tCharCountMap := {
			'A': 1,
			'B': 1,
			'C': 1,
		}
	*/
	tCharCountMap := make(map[byte]int) // key: character, value: count
	for i := 0; i < len(t); i++ {
		tCharCountMap[t[i]]++
	}

	// use a map to store the characters of s in the current window
	wCharCountMap := make(map[byte]int) // key: character, value: count

	// initial sliding window for scanning s
	left, right := 0, 0
	// initial sliding window for result
	start, end := 0, 0

	minLength := len(s) + 1 // * this is the key point, initial minLength with len(s) + 1
	count := len(t)

	// use two pointers, left and right, to represent the current window
	for right < len(s) {
		wCharCountMap[s[right]]++
		// * this is the key point, check if the current character is in tCharCountMap
		if tCharCountMap[s[right]] > 0 && tCharCountMap[s[right]] >= wCharCountMap[s[right]] {
			count--
		}

		// if count == 0, it means the current window contains all the characters of t
		for count == 0 {
			// update the minLength
			if right-left+1 < minLength {
				minLength = right - left + 1
				// update the start and end index of the current window for result
				start = left
				end = right
			}

			// remove the character from the left of the window
			wCharCountMap[s[left]]--
			// * this is the key point, check if the current character is in tCharCountMap
			if tCharCountMap[s[left]] > 0 && tCharCountMap[s[left]] > wCharCountMap[s[left]] {
				count++
			}

			// move window to right
			left++
		}

		// move window to right
		right++
	}

	// if minLength is not updated, it means there is no window contains all the characters of t
	if minLength == len(s)+1 {
		return ""
	}

	// return the result
	return s[start : end+1]
}

func Test_minWindow1(t *testing.T) {
	type args struct {
		s string
		t string
	}
	type expected struct {
		result string
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
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			expected: expected{
				result: "BANC",
			},
		},
		{
			name: "2",
			args: args{
				s: "a",
				t: "a",
			},
			expected: expected{
				result: "a",
			},
		},
		{
			name: "3",
			args: args{
				s: "a",
				t: "aa",
			},
			expected: expected{
				result: "",
			},
		},
		{
			name: "4",
			args: args{
				s: "cabwefgewcwaefgcf",
				t: "cae",
			},
			expected: expected{
				result: "cwae",
			},
		},
		{
			name: "5",
			args: args{
				s: "a",
				t: "b",
			},
			expected: expected{
				result: "",
			},
		},
		{
			name: "6",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			expected: expected{
				result: "BANC",
			},
		},
		{
			name: "7",
			args: args{
				s: "aab",
				t: "aab",
			},
			expected: expected{
				result: "aab",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			minWindow1(tc.args.s, tc.args.t),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
