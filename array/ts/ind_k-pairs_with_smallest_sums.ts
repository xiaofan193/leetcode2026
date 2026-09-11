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

    type Node = { sum: number; i: number; j: number };

    // 先比和，和相同再比下标，保证同和时的输出顺序稳定
    const less = (x: Node, y: Node): boolean =>
        x.sum !== y.sum ? x.sum < y.sum : x.i !== y.i ? x.i < y.i : x.j < y.j;

    // TS 没有内置优先队列，就地写一个最小堆
    const heap: Node[] = [];

    const push = (node: Node): void => {
        heap.push(node);
        // 上浮
        let c = heap.length - 1;
        while (c > 0) {
            const p = (c - 1) >> 1;
            if (!less(heap[c], heap[p])) {
                break;
            }
            [heap[c], heap[p]] = [heap[p], heap[c]];
            c = p;
        }
    };

    const pop = (): Node => {
        const top = heap[0];
        const last = heap.pop() as Node;
        // 下沉
        if (heap.length > 0) {
            heap[0] = last;
            let p = 0;
            for (;;) {
                const l = p * 2 + 1;
                const r = l + 1;
                let s = p;
                if (l < heap.length && less(heap[l], heap[s])) {
                    s = l;
                }
                if (r < heap.length && less(heap[r], heap[s])) {
                    s = r;
                }
                if (s === p) {
                    break;
                }
                [heap[p], heap[s]] = [heap[s], heap[p]];
                p = s;
            }
        }
        return top;
    };

    push({ sum: a[0] + b[0], i: 0, j: 0 });

    while (ans.length < k && heap.length > 0) {
        const cur = pop();
        ans.push([a[cur.i], b[cur.j]]);

        // 同一行向右扩展
        if (cur.j + 1 < b.length) {
            push({ sum: a[cur.i] + b[cur.j + 1], i: cur.i, j: cur.j + 1 });
        }

        // 只有第一列才向下扩展；否则 (i,j) 会被 (i-1,j) 和 (i,j-1) 重复入堆
        if (cur.j === 0 && cur.i + 1 < a.length) {
            push({ sum: a[cur.i + 1] + b[0], i: cur.i + 1, j: 0 });
        }
    }

    // 对数不足 k 时堆会先空，这里自然返回全部，不会越界
    return ans;
}
