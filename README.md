# LeetCode 刷题笔记（2026）

我的 LeetCode 算法刷题记录与总结：以 **Go / Python / TypeScript** 三种语言实现，按算法专题组织整理。

## 📖 目录

| 专题 | 说明 | 链接 |
| ---- | ---- | ---- |
| 📦 动态规划 | 基础 / 0-1 背包 / 股票 / 子序列 | [dp](./dp) |
| 🔁 回溯算法 | 组合 / 分割 / 子集 / 排列 / 棋盘类问题 | [backtracking](./backtracking) |
| 📐 数组 | 双指针 / 原地修改 / 排序后处理 | [array](./array) |

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

数组题围绕**下标的移动与覆盖**展开，常见套路是**双指针（快慢 / 左右 / 荷兰国旗）、原地修改、先排序再复用有序性**，目标多是 O(1) 额外空间。

对应目录：[`array/`](./array)，内含 `golang/`、`python/`、`ts/` 三个语言版本（10 题全部三语齐备）。

### 基础 · 快慢双指针（原地覆盖）

一个「慢指针」标记写入位，「快指针」负责扫描，即可在不新建数组的前提下完成去重 / 移除 / 归位。

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 26 | 删除有序数组中的重复项 | [remove_duplicates_from_sorted_array.go](./array/golang/remove_duplicates_from_sorted_array.go) | [remove_duplicates_from_sorted_array.py](./array/python/remove_duplicates_from_sorted_array.py) | [remove_duplicates_from_sorted_array.ts](./array/ts/remove_duplicates_from_sorted_array.ts) |
| 80 | 删除有序数组中的重复项 II | [remove_duplicates_from_sorted_array_ii.go](./array/golang/remove_duplicates_from_sorted_array_ii.go) | [remove_duplicates_from_sorted_array_ii.py](./array/python/remove_duplicates_from_sorted_array_ii.py) | [remove_duplicates_from_sorted_array_ii.ts](./array/ts/remove_duplicates_from_sorted_array_ii.ts) |
| 27 | 移除元素 | [remove_element.go](./array/golang/remove_element.go) | [remove_element.py](./array/python/remove_element.py) | [remove_element.ts](./array/ts/remove_element.ts) |
| 283 | 移动零 | [move_zeroes.go](./array/golang/move_zeroes.go) | [move_zeroes.py](./array/python/move_zeroes.py) | [move_zeros.ts](./array/ts/move_zeros.ts) |

> 💡 **26 → 80** 是同一模板的进阶：去重条件从「每个元素保留 1 次」放宽到「最多保留 2 次」，指针的写入 / 比较时机随之变化。

### 进阶 · 已排序 / 排序后处理

输入本就有序（88 / 1287）或先 `sort` 的题目，再配合**双指针、相邻比较、区间统计**，往往能把组合遍历退化为线性扫描。

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 88 | 合并两个有序数组 | [merge_sorted_array.go](./array/golang/merge_sorted_array.go) | [merge_sorted_array.py](./array/python/merge_sorted_array.py) | [merge_sorted_array.ts](./array/ts/merge_sorted_array.ts) |
| 349 | 两个数组的交集 | [ntersection_of_two_arrays.go](./array/golang/ntersection_of_two_arrays.go) | [ntersection-of-two-arrays.py](./array/python/ntersection-of-two-arrays.py) | [ntersection-of-two-arrays.ts](./array/ts/ntersection-of-two-arrays.ts) |
| 1200 | 最小绝对差 | [minimum_absolute_difference.go](./array/golang/minimum_absolute_difference.go) | [minimum_absolute_difference.py](./array/python/minimum_absolute_difference.py) | [minimum_absolute_difference.ts](./array/ts/minimum_absolute_difference.ts) |
| 1287 | 有序数组中出现次数超过 25% 的元素 | [element_appearing_more_than-_5-_n_sorted_array.go](./array/golang/element_appearing_more_than-_5-_n_sorted_array.go) | [element_appearing_ore_than_25_in_sorted_array.py](./array/python/element_appearing_ore_than_25_in_sorted_array.py) | [element_appearing_ore_than_25_in_sorted_array.ts](./array/ts/element_appearing_ore_than_25_in_sorted_array.ts) |

> 💡 **88** 用「从后往前」归并，避免覆盖 `nums1` 尚未处理的元素；**349** 先排序再双指针并**跳过重复值**，保证结果唯一。
>
> 📎 88 另附 Go 单元测试：[merge_sorted_array_test.go](./array/golang/merge_sorted_array_test.go)

### 综合 · 分区、翻转与轮转

| 题号 | 题目 | Go | Python | TypeScript |
| ---- | ---- | :-: | :-: | :-: |
| 75 | 颜色分类 | [sort_color.go](./array/golang/sort_color.go) | [sort_color.py](./array/python/sort_color.py) | [sort_color.ts](./array/ts/sort_color.ts) |
| 189 | 轮转数组 | [rotate_array.go](./array/golang/rotate_array.go) | [rotae_array.py](./array/python/rotae_array.py) | [rotate_array.ts](./array/ts/rotate_array.ts) |

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
├── array/               # 数组
│   ├── golang/          # Go 实现
│   ├── python/          # Python 实现
│   └── ts/              # TypeScript 实现
├── main.go
├── main.py
├── go.mod
└── README.md
```

## 🛠 环境

- Go 1.26
- Python 3
- TypeScript / Node.js
