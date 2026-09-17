package golang

import "container/heap"

// 1439. 有序矩阵中的第 k 个最小数组和

//https://leetcode.cn/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/

// 给你一个 m * n 的矩阵 mat，以及一个整数 k ，矩阵中的每一行都以非递减的顺序排列。

// 你可以从每一行中选出 1 个元素形成一个数组。返回所有可能数组中的第 k 个 最小 数组和。

// 示例 1：

// 输入：mat = [[1,3,11],[2,4,6]], k = 5
// 输出：7
// 解释：从每一行中选出一个元素，前 k 个和最小的数组分别是：
// [1,2], [1,4], [3,2], [3,4], [1,6]。其中第 5 个的和是 7 。
// 示例 2：

// 输入：mat = [[1,3,11],[2,4,6]], k = 9
// 输出：17
// 示例 3：

// 输入：mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
// 输出：9
// 解释：从每一行中选出一个元素，前 k 个和最小的数组分别是：
// [1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]。其中第 7 个的和是 9 。
// 示例 4：

// 输入：mat = [[1,1,10],[2,2,9]], k = 7
// 输出：12

// 思路：逐行合并，每一轮只保留前 k 小的和。
//
// 1) 把 m 行拆成两两合并：prev 表示"前 i 行"能凑出的前 k 小和（升序，长度 ≤ k），
// 再拿 prev 和第 i+1 行做 merge，得到"前 i+1 行"的前 k 小和，逐行滚下去。
// 2) mergeTopKSum(f, g, k) 就是「两个有序数组取前 k 小的两两和」，等价于 373 题：
// 把 f[i] + g[j] 看成 i×j 的矩阵，行、列都递增，用堆维护每个 j 对应的那一条链上的
// 当前最小值；弹出堆顶后沿着 f 方向（f 下标 +1）继续推进。因为 f 升序，同一条链上的
// 和也递增，所以每次弹堆顶就是全局下一个最小值，不会漏也不会重。
// 3) 每轮截断到 k 不影响正确性：若"前 i 行"的某个中间和 s 排不进前 k，那已经存在 k 个
// 比 s 更小的中间和；它们后续都选同样的元素，得到的完整和也都比 s 小，因此 s 不可能
// 贡献最终的前 k 小。
//
// 复杂度：每轮 mergeTopKSum 是 O(n + k log n)（n 为较短一侧的长度），共 m - 1 轮，
// 整体 O(m * k log n)，空间 O(k + n)。
func kthSmallest(mat [][]int, k int) int {
	// prev 始终是"前几行"的前 k 小和且保持升序；mat[0] 本身升序，满足合并的前提
	prev := mat[0]
	for i := 1; i < len(mat); i++ {
		prev = mergeTopKSum(prev, mat[i], k)
	}
	// prev 升序，第 k 小（1-based）即下标 k-1
	return prev[k-1]
}

// mergeTopKSum 返回 f、g 两两相加后的前 k 小和（升序）。前提：f、g 均升序。
// 复用 373 题的 pairItem/pairHeap：i 表示 f 的下标，j 表示 g 的下标，sum 为两者之和。
func mergeTopKSum(f, g []int, k int) []int {
	// 堆里每条链对应一个 g 下标，链数 = len(g)，所以让 f 是较长的一边，
	// 堆的规模就等于较短一边的长度（只是常数优化，对结果无影响）
	if len(g) > len(f) {
		f, g = g, f
	}

	// 每条链 j 的起点都是 (f[0] + g[j])
	h := make(pairHeap, len(g))
	for j := range g {
		h[j] = pairItem{sum: f[0] + g[j], i: 0, j: j}
	}
	heap.Init(&h)

	ans := make([]int, 0, k)
	for k > 0 && h.Len() > 0 {
		cur := heap.Pop(&h).(pairItem)
		ans = append(ans, cur.sum)

		// 沿 f 方向推进：g 的下标不变，f 的下标 +1，即同一条链上的下一项
		if cur.i+1 < len(f) {
			heap.Push(&h, pairItem{sum: f[cur.i+1] + g[cur.j], i: cur.i + 1, j: cur.j})
		}
		k--
	}

	// 组合数不足 k 时堆会先空，这里自然返回全部，不会越界
	return ans
}
