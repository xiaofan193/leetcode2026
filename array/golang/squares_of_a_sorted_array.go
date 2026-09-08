package golang

// 977. 有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/description/
// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

// 示例 1：
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 示例 2：
// 输入：nums = [-7,-3,2,3,11]
// 输出：[4,9,9,49,121]

// 双指针：原数组升序，平方后的最大值只可能来自「绝对值最大的两端」，
// 即最左的负数或最右的正数。i、j 分别指向首尾，比较平方大小，从结果数组末尾往前填。
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	i, j := 0, n-1
	for k := n - 1; k >= 0; k-- {
		a, b := nums[i]*nums[i], nums[j]*nums[j]
		if a > b {
			res[k] = a
			i++
		} else {
			res[k] = b
			j--
		}
	}
	return res
}
