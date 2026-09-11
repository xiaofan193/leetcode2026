package golang

// 986. 区间列表的交集

// https://leetcode.cn/problems/interval-list-intersections/description/

// 给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。每个区间列表都是成对 不相交 的，并且 已经排序 。

// 返回这 两个区间列表的交集 。

// 形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。

// 两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。

// 示例 1：

// 输入：firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
// 输出：[[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
// 示例 2：

// 输入：firstList = [[1,3],[5,9]], secondList = []
// 输出：[]
// 示例 3：

// 输入：firstList = [], secondList = [[4,8],[10,12]]
// 输出：[]
// 示例 4：

// 输入：firstList = [[1,7]], secondList = [[3,10]]
// 输出：[[3,7]]

func maxInt(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func minInt(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func intervalIntersection(firstList [][]int, secondList [][]int) (ans [][]int) {
	// 1.分别定义指向firstList和secondList的双指针
	p, q := 0, 0
	for p < len(firstList) && q < len(secondList) {
		// 2.找寻p,q指向的区间的交集
		left := maxInt(firstList[p][0], secondList[q][0])
		right := minInt(firstList[p][1], secondList[q][1])
		// 如果left<=right，则说明两个区间存在交集,将其添加进ans
		if left <= right {
			ans = append(ans, []int{left, right})
		}
		/* 3.指针右移:
		(1)若p指向的右端点<=q指向的右端点,则p++
		(2)若p指向的右端点>q指向的右端点,则q++
		*/
		if firstList[p][1] <= secondList[q][1] {
			p++
		} else {
			q++
		}
	}
	return
}
