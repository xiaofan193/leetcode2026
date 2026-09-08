package golang

// element_appearing_ore_than_25_in_sorted_array

// 1287 有序数组出现次数超过元素总数25%的元素

// 给你一个非递减的 有序 整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。

// 请你找到并返回这个整数

// 示例：

// 输入：arr = [1,2,2,6,6,6,6,7,10]
// 输出：6

func findSpecialInteger(arr []int) int {
	c, k := arr[0], 0
	// c 为出现次数为k的字符
	// k * 4 > n

	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] == c {
			k++
			if k*4 > n {
				return c
			}
		} else {
			c, k = arr[i], 1
		}
	}
	return c
}
