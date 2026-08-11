# LeetCode 刷题笔记（2026）

我的 LeetCode 算法刷题记录与总结：以 **Go / Python / TypeScript** 三种语言实现，按算法专题组织整理。

## 📖 目录

| 专题 | 说明 | 链接 |
| ---- | ---- | ---- |
| 🔁 回溯算法 | 组合 / 分割 / 子集 / 排列 / 棋盘类问题 | [backtracking](./backtracking) |

> 持续更新中，后续将补充动态规划、贪心、二分查找等专题。

## 🔁 回溯算法

回溯算法（Backtracking）通过「递归 + 回溯」遍历所有可能的解空间，并借助剪枝优化，常用于解决**组合、切割、子集、排列、棋盘**等「暴力枚举」类问题。

对应目录：[`backtracking/`](./backtracking)，内含 `golang/`、`python/`、`ts/` 三个语言版本。

### 组合问题

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 77 | 组合 | — | [combinations.py](./backtracking/python/combinations.py) | [combinations.ts](./backtracking/ts/combinations.ts) |
| 216 | 组合总和 III | [combinations.go](./backtracking/golang/combinations.go) | — | — |
| 39 | 组合总和 | [combinationSum.go](./backtracking/golang/combinationSum.go) | — | — |
| 40 | 组合总和 II | — | [combinationSum2.py](./backtracking/python/combinationSum2.py) | [combinationSum2.ts](./backtracking/ts/combinationSum2.ts) |

### 分割问题

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 131 | 分割回文串 | [palindrome_partitioning.go](./backtracking/golang/palindrome_partitioning.go) | [palindrome_partitioning.py](./backtracking/python/palindrome_partitioning.py) | [palindrome_partitioning.ts](./backtracking/ts/palindrome_partitioning.ts) |
| 93 | 复原 IP 地址 | [restoreIp_addresses.go](./backtracking/golang/restoreIp_addresses.go) | [restoreIp_addresses.py](./backtracking/python/restoreIp_addresses.py) | [restoreIp_addresses.ts](./backtracking/ts/restoreIp_addresses.ts) |

### 子集问题

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 78 | 子集 | [subsets.go](./backtracking/golang/subsets.go) | [subsets.py](./backtracking/python/subsets.py) | [subsets.ts](./backtracking/ts/subsets.ts) |
| 491 | 非递减子序列 | [non_decreasing_subsequences.go](./backtracking/golang/non_decreasing_subsequences.go) | [non_decreasing_subsequences.py](./backtracking/python/non_decreasing_subsequences.py) | [non_decreasing_subsequences.ts](./backtracking/ts/non_decreasing_subsequences.ts) |

### 排列问题

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 46 | 全排列 | [permutations.go](./backtracking/golang/permutations.go) | [permutations.py](./backtracking/python/permutations.py) | [permutations.ts](./backtracking/ts/permutations.ts) |

### 其他

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 17 | 电话号码的字母组合 | [letterCombinations.go](./backtracking/golang/letterCombinations.go) | — | [letterCombinations.ts](./backtracking/ts/letterCombinations.ts) |
| 332 / 51 / 37 | 重新安排行程 / N 皇后 / 解数独 | [hard.go](./backtracking/golang/hard.go) | — | — |

### 回溯算法模板

```python
def backtracking(参数):
    if 终止条件:
        存放结果
        return

    for 选择 in 本层可选择的元素:
        处理节点
        backtracking(路径, 选择列表)   # 递归
        回溯，撤销处理结果
```

## 📁 目录结构

```
leetcode2026/
├── backtracking/        # 回溯算法
│   ├── golang/          # Go 实现
│   ├── python/          # Python 实现
│   └── ts/              # TypeScript 实现
├── main.go
├── go.mod
└── README.md
```

## 🛠 环境

- Go 1.26
- Python 3
- TypeScript / Node.js
