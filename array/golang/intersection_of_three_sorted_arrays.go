package golang

// 1213 3个有序数组的交集
// intersection_—of_three_sorted_arrays

// https://leetcode.cn/problems/intersection-of-three-sorted-arrays/description/

// 给定3个严格递增的排列的整数数组 arr1 arr2 arr3,返回一个仅由这3个数组中同时出现整数构成的有序数组
// 例如： arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9],arr3 = [1,3,4,5,8]
// 答案: [1,5]
func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	res := make([]int, 0)

	// 三指针同向扫描：三个指针都指向当前最小的未比较元素
	i, j, k := 0, 0, 0

	for i < len(arr1) && j < len(arr2) && k < len(arr3) {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			// 数组严格递增，相等元素唯一，无需去重
			res = append(res, arr1[i])
			i, j, k = i+1, j+1, k+1
		} else {
			// 把三者中最小的指针右移，向更大的值靠拢
			min := arr1[i]
			if arr2[j] < min {
				min = arr2[j]
			}
			if arr3[k] < min {
				min = arr3[k]
			}
			if arr1[i] == min {
				i++
			}
			if arr2[j] == min {
				j++
			}
			if arr3[k] == min {
				k++
			}
		}
	}

	return res
}
