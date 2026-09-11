// find_k-pairs_with_smallest_sums
// 373. 查找和最小的 K 对数字
// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
//
// 给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
// 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
//
// 示例 1:
//
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [[1,2],[1,4],[1,6]]
// 解释: 返回序列中的前 3 对数：
//      [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
// 示例 2:
//
// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [[1,1],[1,1]]

type PairNode = { sum: number; i: number; j: number };

// 先比和，和相同再比下标，保证同和时的输出顺序稳定
function pairLess(a: PairNode, b: PairNode): boolean {
    if (a.sum !== b.sum) {
        return a.sum < b.sum;
    }
    if (a.i !== b.i) {
        return a.i < b.i;
    }
    return a.j < b.j;
}

// TS 没有内置优先队列，手写一个最小堆
class PairMinHeap {
    private data: PairNode[] = [];

    get size(): number {
        return this.data.length;
    }

    push(node: PairNode): void {
        const d = this.data;
        d.push(node);

        // 上浮
        let c = d.length - 1;
        while (c > 0) {
            const p = (c - 1) >> 1;
            if (!pairLess(d[c], d[p])) {
                break;
            }
            [d[c], d[p]] = [d[p], d[c]];
            c = p;
        }
    }

    pop(): PairNode {
        const d = this.data;
        const top = d[0];
        const last = d.pop() as PairNode;

        // 下沉
        if (d.length > 0) {
            d[0] = last;
            let p = 0;
            for (;;) {
                const l = p * 2 + 1;
                const r = l + 1;
                let s = p;
                if (l < d.length && pairLess(d[l], d[s])) {
                    s = l;
                }
                if (r < d.length && pairLess(d[r], d[s])) {
                    s = r;
                }
                if (s === p) {
                    break;
                }
                [d[p], d[s]] = [d[s], d[p]];
                p = s;
            }
        }

        return top;
    }
}

// 思路：把数对看成 m×n 的矩阵，行列都递增，问题变成"在有序矩阵里取前 k 小"。
// 用堆维护"已发现的候选边界"：初始只有 (0,0)，每次弹出堆顶后再把它的右邻居和
// 下邻居推进堆。为避免同一个格子被两条路径重复入堆，只让第一列向下扩展。
function kSmallestPairs(nums1: number[], nums2: number[], k: number): number[][] {
    // sort 原地排序，先拷贝一份，避免改动调用方传入的数组
    const a = [...nums1].sort((x, y) => x - y);
    const b = [...nums2].sort((x, y) => x - y);

    const ans: number[][] = [];
    if (a.length === 0 || b.length === 0 || k <= 0) {
        return ans;
    }

    const heap = new PairMinHeap();
    heap.push({ sum: a[0] + b[0], i: 0, j: 0 });

    while (ans.length < k && heap.size > 0) {
        const cur = heap.pop();
        ans.push([a[cur.i], b[cur.j]]);

        // 同一行向右扩展
        if (cur.j + 1 < b.length) {
            heap.push({ sum: a[cur.i] + b[cur.j + 1], i: cur.i, j: cur.j + 1 });
        }

        // 只有第一列才向下扩展；否则 (i,j) 会被 (i-1,j) 和 (i,j-1) 重复入堆
        if (cur.j === 0 && cur.i + 1 < a.length) {
            heap.push({ sum: a[cur.i + 1] + b[0], i: cur.i + 1, j: 0 });
        }
    }

    // 对数不足 k 时堆会先空，这里自然返回全部，不会越界
    return ans;
}
