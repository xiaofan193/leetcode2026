package golang

// 360 有序数组转化
// https://leetcode.cn/problems/sort-transformed-array/

// Given a sorted array of integers nums and integer values a, band c. Apply a quadratic function of the form f(x) = ax2 + bx + c to each element x in the array.

// The returned array must be in sorted order.

// Expected time complexity: O(n)

// Example 1:

// Input: nums = [-4,-2,2,4], a = 1, b = 3, c = 5
// Output: [3,9,15,33]
// Example 2:

// Input: nums = [-4,-2,2,4], a = -1, b = 3, c = 5
// Output: [-23,-5,1,7]

func sortTransformedArray(nums []int, a, b, c int) []int {
	n := len(nums)

	left, right := 0, n-1
	k := 0
	if a >= 0 {
		k = n - 1
	}
	res := make([]int, n)

	for left <= right {
		left_val := quadraticValue(nums[left], a, b, c)
		right_val := quadraticValue(nums[right], a, b, c)

		// 二次函数开口向上 两边增大
		if a >= 0 {
			if left_val >= right_val {
				res[k] = left_val
				left++
			} else {
				res[k] = right_val
				right--
			}
			k--
		} else {
			if left_val <= right_val {
				res[k] = left_val
				left++
			} else {
				res[k] = right_val
				right--
			}
			k++
		}
	}
	return res
}

func quadraticValue(x, a, b, c int) int {
	return a*x*x + b*x + c
}
