package golang

import "sort"

// ntersection-of-two-arrays
// 349 两个数字的交集
// 给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

// 示例 1：

// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]
// 示例 2：

// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[9,4]
// 解释：[4,9] 也是可通过的
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	m, n := len(nums1), len(nums2)

	i, j := 0, 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			// 排序后重复的相等元素必然连续，和 res 末位相同即为重复，跳过
			if len(res) == 0 || res[len(res)-1] != nums1[i] {
				res = append(res, nums1[i])
			}
			i, j = i+1, j+1
		}

	}
	return res
}
