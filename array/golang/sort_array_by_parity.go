package golang

// 905 按奇偶排序数组
// https://leetcode.cn/problems/sort-array-by-parity/

func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j { // 循环直到不足两个数
		if nums[i]%2 == 0 { // 寻找最左边的奇数
			i++
		} else if nums[j]%2 == 1 { // 寻找最右边的偶数
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			// 交换后，问题变成 [i+1,j-1] 的子问题
			i++
			j--
		}
	}
	return nums

}
