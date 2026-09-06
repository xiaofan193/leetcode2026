# LeetCode 刷题笔记（2026）

我的 LeetCode 算法刷题记录与总结：以 **Go / Python / TypeScript** 三种语言实现，按算法专题组织整理。

## 📖 目录

| 专题 | 说明 | 链接 |
| ---- | ---- | ---- |
| 📦 动态规划 | 基础 / 0-1 背包 / 股票 / 子序列 | [dp](./dp) |
| 🔁 回溯算法 | 组合 / 分割 / 子集 / 排列 / 棋盘类问题 | [backtracking](./backtracking) |
| 📐 数组 | 数组遍历 / 双指针 / 滑动窗口（规划中） | [array](./array) |

> 持续更新中，后续将补充贪心、二分查找、链表、二叉树等专题。

---

## 📦 动态规划（dp）

动态规划（Dynamic Programming）通过**分解子问题**并**保存子问题的解**（通常为 dp 数组/表格）来避免重复计算，适用于具备「最优子结构」与「重叠子问题」特性的问题，核心在于确定 **dp 数组含义 → 递推公式 → 初始化 → 遍历顺序**。

对应目录：[`dp/`](./dp)，内含 `golang/`、`python/`、`ts/` 三个语言版本。

### 基础题目

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 509 | 斐波那契数 | [fibonacci_number.go](./dp/golang/fibonacci_number.go) | [fibonacci_number.py](./dp/python/fibonacci_number.py) | [fibonacci_number.ts](./dp/ts/fibonacci_number.ts) |
| 70 | 爬楼梯 | [climbing_stairs.go](./dp/golang/climbing_stairs.go) | [climbing_stairs.py](./dp/python/climbing_stairs.py) | [climbing_stairs.ts](./dp/ts/climbing_stairs.ts) |
| 746 | 使用最小花费爬楼梯 | [min_cost_climbing_stairs.go](./dp/golang/min_cost_climbing_stairs.go) | [min_cost_climbing_stairs.py](./dp/python/min_cost_climbing_stairs.py) | [min_cost_climbing_stairs.ts](./dp/ts/min_cost_climbing_stairs.ts) |
| 62 | 不同路径 | [unique_path.go](./dp/golang/unique_path.go) | [unique_path.py](./dp/python/unique_path.py) | [unique_path.ts](./dp/ts/unique_path.ts) |
| 63 | 不同路径 II | [unique_path_ii.go](./dp/golang/unique_path_ii.go) | [unique_path_ii.py](./dp/python/unique_path_ii.py) | [unique_path_ii.ts](./dp/ts/unique_path_ii.ts) |
| 343 | 整数拆分 | [integer_break.go](./dp/golang/integer_break.go) | [integer_break.py](./dp/python/integer_break.py) | [integer_break.ts](./dp/ts/integer_break.ts) |
| 96 | 不同的二叉搜索树 | [unique_binary_search_trees.go](./dp/golang/unique_binary_search_trees.go) | [unique_binary_search_trees.py](./dp/python/unique_binary_search_trees.py) | [unique_nary_search_trees.ts](./dp/ts/unique_nary_search_trees.ts) |

### 0-1 背包问题

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| — | 0-1 背包（二维 dp） | [zero-one-package.go](./dp/golang/zero-one-package.go) | — | — |
| 416 | 分割等和子集 | [partition_equal_subset_sum.go](./dp/golang/partition_equal_subset_sum.go) | [partition_equal_subset_sum.py](./dp/python/partition_equal_subset_sum.py) | [partition_equal_subset_sum.ts](./dp/ts/partition_equal_subset_sum.ts) |
| 1049 | 最后一块石头的重量 II | [last_stone_weight_ii.go](./dp/golang/last_stone_weight_ii.go) | [last_stone_weight_ii.py](./dp/python/last_stone_weight_ii.py) | [last_stone_weight_ii.ts](./dp/ts/last_stone_weight_ii.ts) |
| 494 | 目标和 | [target_sum.go](./dp/golang/target_sum.go) | [target_sum.py](./dp/python/target_sum.py) | [target_sum.ts](./dp/ts/target_sum.ts) |
| 474 | 一和零 | [ones-and-zeroes.go](./dp/golang/ones-and-zeroes.go) | [ones-and-zeroes.py](./dp/python/ones-and-zeroes.py) | [ones-and-zeroes.ts](./dp/ts/ones-and-zeroes.ts) |

### 股票问题

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 121 | 买卖股票的最佳时机 | [best_time_to_buy_and_sell_stock.go](./dp/golang/best_time_to_buy_and_sell_stock.go) | [best_time_to_buy_and_sell_stock.py](./dp/python/best_time_to_buy_and_sell_stock.py) | [best_time_to_buy_and_sell_stock.ts](./dp/ts/best_time_to_buy_and_sell_stock.ts) |
| 122 | 买卖股票的最佳时机 II | [best_time_to_buy_and_sell_stock_ii.go](./dp/golang/best_time_to_buy_and_sell_stock_ii.go) | — | — |

### 子序列问题

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 516 | 最长回文子序列 | [longest_palindromic_subsequence.go](./dp/golang/longest_palindromic_subsequence.go) | [longest_palindromic_subsequence.py](./dp/python/longest_palindromic_subsequence.py) | [longest_palindromic_subsequence.ts](./dp/ts/longest_palindromic_subsequence.ts) |

> 122 题另附 Go 单元测试：[best_time_to_buy_and_sell_stock_ii_test.go](./dp/golang/best_time_to_buy_and_sell_stock_ii_test.go)

---

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

---

## 📐 数组（array）

数组是算法题中最基础的数据结构，围绕它的解题技巧通常包括**双指针、滑动窗口、前缀和、二分查找、螺旋矩阵（循环不变量）**等。

对应目录：[`array/`](./array)。

> ⏳ 该专题正在整理中，问题会以 `array/golang/`、`array/python/`、`array/ts/` 三个子目录补充，敬请期待。

---

## 📁 目录结构

```
leetcode2026/
├── backtracking/        # 回溯算法
│   ├── golang/          # Go 实现
│   ├── python/          # Python 实现
│   └── ts/              # TypeScript 实现
├── dp/                  # 动态规划
│   ├── golang/          # Go 实现
│   ├── python/          # Python 实现
│   └── ts/              # TypeScript 实现
├── array/               # 数组（待补充）
├── main.go
├── go.mod
└── README.md
```

## 🛠 环境

- Go 1.26
- Python 3
- TypeScript / Node.js
