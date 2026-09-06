package golang

// 88 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/
// 思路：nums1 长度已是 m+n，末尾 n 个为占位 0，需原地合并（不能新建切片返回）。
// 从后往前比较并填充，避免了从前往后归并会覆盖 nums1 前部未处理元素的问题。
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	// j >= 0 直接表达"只需要一直归并到 nums2 清空为止",终止条件更自然,也不用处理越界
	for k := m + n - 1; j >= 0; k-- {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
