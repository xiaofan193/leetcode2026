package golang

import (
	"container/heap"
	"sort"
)

// find_k-pairs_with_smallest_sums
// 373. 查找和最小的 K 对数字
// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/

// 给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

// 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

// 示例 1:

// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [[1,2],[1,4],[1,6]]
// 解释: 返回序列中的前 3 对数：
//      [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
// 示例 2:

// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [[1,1],[1,1]]
// 解释: 返回序列中的前 2 对数：
//      [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

// pairItem 是一个候选数对：(i, j) 表示取 nums1[i] 和 nums2[j]
type pairItem struct {
	sum  int
	i, j int
}

// pairHeap 小顶堆，堆顶始终是当前候选里和最小的一对
type pairHeap []pairItem

func (h pairHeap) Len() int { return len(h) }

// 先比和，和相同再比下标，保证同和时的输出顺序稳定
func (h pairHeap) Less(a, b int) bool {
	if h[a].sum != h[b].sum {
		return h[a].sum < h[b].sum
	}
	if h[a].i != h[b].i {
		return h[a].i < h[b].i
	}
	return h[a].j < h[b].j
}

func (h pairHeap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *pairHeap) Push(x any) { *h = append(*h, x.(pairItem)) }

func (h *pairHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// 思路：把数对看成 m×n 的矩阵，行列都递增，问题变成"在有序矩阵里取前 k 小"。
// 用堆维护"已发现的候选边界"：初始只有 (0,0)，每次弹出堆顶后再把它的右邻居和
// 下邻居推进堆。为避免同一个格子被两条路径重复入堆，只让第一列向下扩展。
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return [][]int{}
	}

	h := &pairHeap{{sum: nums1[0] + nums2[0], i: 0, j: 0}}
	heap.Init(h)

	ans := make([][]int, 0, min(k, len(nums1)*len(nums2)))
	for len(ans) < k && h.Len() > 0 {
		cur := heap.Pop(h).(pairItem)
		ans = append(ans, []int{nums1[cur.i], nums2[cur.j]})

		// 同一行向右扩展
		if cur.j+1 < len(nums2) {
			heap.Push(h, pairItem{sum: nums1[cur.i] + nums2[cur.j+1], i: cur.i, j: cur.j + 1})
		}

		// 只有第一列才向下扩展；否则 (i,j) 会被 (i-1,j) 和 (i,j-1) 重复入堆
		if cur.j == 0 && cur.i+1 < len(nums1) {
			heap.Push(h, pairItem{sum: nums1[cur.i+1] + nums2[0], i: cur.i + 1, j: 0})
		}
	}

	// 对数不足 k 时堆会先空，这里自然返回全部，不会越界
	return ans
}
