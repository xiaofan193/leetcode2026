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

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 509 | 斐波那契数 | [✅](./dp/golang/fibonacci_number.go) | [✅](./dp/python/fibonacci_number.py) | [✅](./dp/ts/fibonacci_number.ts) |
| 70 | 爬楼梯 | [✅](./dp/golang/climbing_stairs.go) | [✅](./dp/python/climbing_stairs.py) | [✅](./dp/ts/climbing_stairs.ts) |
| 746 | 使用最小花费爬楼梯 | [✅](./dp/golang/min_cost_climbing_stairs.go) | [✅](./dp/python/min_cost_climbing_stairs.py) | [✅](./dp/ts/min_cost_climbing_stairs.ts) |
| 62 | 不同路径 | [✅](./dp/golang/unique_path.go) | [✅](./dp/python/unique_path.py) | [✅](./dp/ts/unique_path.ts) |
| 63 | 不同路径 II | [✅](./dp/golang/unique_path_ii.go) | [✅](./dp/python/unique_path_ii.py) | [✅](./dp/ts/unique_path_ii.ts) |
| 343 | 整数拆分 | [✅](./dp/golang/integer_break.go) | [✅](./dp/python/integer_break.py) | [✅](./dp/ts/integer_break.ts) |
| 96 | 不同的二叉搜索树 | [✅](./dp/golang/unique_binary_search_trees.go) | [✅](./dp/python/unique_binary_search_trees.py) | [✅](./dp/ts/unique_nary_search_trees.ts) |

### 0-1 背包问题

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| — | 0-1 背包（二维 dp） | [✅](./dp/golang/zero-one-package.go) | — | — |
| 416 | 分割等和子集 | [✅](./dp/golang/partition_equal_subset_sum.go) | [✅](./dp/python/partition_equal_subset_sum.py) | [✅](./dp/ts/partition_equal_subset_sum.ts) |
| 1049 | 最后一块石头的重量 II | [✅](./dp/golang/last_stone_weight_ii.go) | [✅](./dp/python/last_stone_weight_ii.py) | [✅](./dp/ts/last_stone_weight_ii.ts) |
| 494 | 目标和 | [✅](./dp/golang/target_sum.go) | [✅](./dp/python/target_sum.py) | [✅](./dp/ts/target_sum.ts) |
| 474 | 一和零 | [✅](./dp/golang/ones-and-zeroes.go) | [✅](./dp/python/ones-and-zeroes.py) | [✅](./dp/ts/ones-and-zeroes.ts) |

### 股票问题

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 121 | 买卖股票的最佳时机 | [✅](./dp/golang/best_time_to_buy_and_sell_stock.go) | [✅](./dp/python/best_time_to_buy_and_sell_stock.py) | [✅](./dp/ts/best_time_to_buy_and_sell_stock.ts) |
| 122 | 买卖股票的最佳时机 II | [✅](./dp/golang/best_time_to_buy_and_sell_stock_ii.go) | — | — |

### 子序列问题

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 516 | 最长回文子序列 | [✅](./dp/golang/longest_palindromic_subsequence.go) | [✅](./dp/python/longest_palindromic_subsequence.py) | [✅](./dp/ts/longest_palindromic_subsequence.ts) |

> 122 题另附 Go 单元测试：[测试](./dp/golang/best_time_to_buy_and_sell_stock_ii_test.go)

---

## 🔁 回溯算法

回溯算法（Backtracking）通过「递归 + 回溯」遍历所有可能的解空间，并借助剪枝优化，常用于解决**组合、切割、子集、排列、棋盘**等「暴力枚举」类问题。

对应目录：[`backtracking/`](./backtracking)，内含 `golang/`、`python/`、`ts/` 三个语言版本。

### 组合问题

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 77 | 组合 | — | [✅](./backtracking/python/combinations.py) | [✅](./backtracking/ts/combinations.ts) |
| 216 | 组合总和 III | [✅](./backtracking/golang/combinations.go) | — | — |
| 39 | 组合总和 | [✅](./backtracking/golang/combinationSum.go) | — | — |
| 40 | 组合总和 II | — | [✅](./backtracking/python/combinationSum2.py) | [✅](./backtracking/ts/combinationSum2.ts) |

### 分割问题

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 131 | 分割回文串 | [✅](./backtracking/golang/palindrome_partitioning.go) | [✅](./backtracking/python/palindrome_partitioning.py) | [✅](./backtracking/ts/palindrome_partitioning.ts) |
| 93 | 复原 IP 地址 | [✅](./backtracking/golang/restoreIp_addresses.go) | [✅](./backtracking/python/restoreIp_addresses.py) | [✅](./backtracking/ts/restoreIp_addresses.ts) |

### 子集问题

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 78 | 子集 | [✅](./backtracking/golang/subsets.go) | [✅](./backtracking/python/subsets.py) | [✅](./backtracking/ts/subsets.ts) |
| 491 | 非递减子序列 | [✅](./backtracking/golang/non_decreasing_subsequences.go) | [✅](./backtracking/python/non_decreasing_subsequences.py) | [✅](./backtracking/ts/non_decreasing_subsequences.ts) |

### 排列问题

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 46 | 全排列 | [✅](./backtracking/golang/permutations.go) | [✅](./backtracking/python/permutations.py) | [✅](./backtracking/ts/permutations.ts) |

### 其他

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 17 | 电话号码的字母组合 | [✅](./backtracking/golang/letterCombinations.go) | — | [✅](./backtracking/ts/letterCombinations.ts) |
| 332 / 51 / 37 | 重新安排行程 / N 皇后 / 解数独 | [✅](./backtracking/golang/hard.go) | — | — |

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

对应目录：[`array/`](./array)，内含 `golang/`、`python/`、`ts/` 三个语言版本（11 题全部三语齐备）。

### 基础 · 快慢双指针（原地覆盖）

一个「慢指针」标记写入位，「快指针」负责扫描，即可在不新建数组的前提下完成去重 / 移除 / 归位。

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 26 | 删除有序数组中的重复项 | [✅](./array/golang/remove_duplicates_from_sorted_array.go) | [✅](./array/python/remove_duplicates_from_sorted_array.py) | [✅](./array/ts/remove_duplicates_from_sorted_array.ts) |
| 80 | 删除有序数组中的重复项 II | [✅](./array/golang/remove_duplicates_from_sorted_array_ii.go) | [✅](./array/python/remove_duplicates_from_sorted_array_ii.py) | [✅](./array/ts/remove_duplicates_from_sorted_array_ii.ts) |
| 27 | 移除元素 | [✅](./array/golang/remove_element.go) | [✅](./array/python/remove_element.py) | [✅](./array/ts/remove_element.ts) |
| 283 | 移动零 | [✅](./array/golang/move_zeroes.go) | [✅](./array/python/move_zeroes.py) | [✅](./array/ts/move_zeros.ts) |

> 💡 **26 → 80** 是同一模板的进阶：去重条件从「每个元素保留 1 次」放宽到「最多保留 2 次」，指针的写入 / 比较时机随之变化。

### 进阶 · 已排序 / 排序后处理

输入本就有序（88 / 1287）或先 `sort` 的题目，再配合**双指针、相邻比较、区间统计**，往往能把组合遍历退化为线性扫描。

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 88 | 合并两个有序数组 | [✅](./array/golang/merge_sorted_array.go) | [✅](./array/python/merge_sorted_array.py) | [✅](./array/ts/merge_sorted_array.ts) |
| 349 | 两个数组的交集 | [✅](./array/golang/ntersection_of_two_arrays.go) | [✅](./array/python/ntersection-of-two-arrays.py) | [✅](./array/ts/ntersection-of-two-arrays.ts) |
| 977 | 有序数组的平方 | [✅](./array/golang/squares_of_a_sorted_array.go) | [✅](./array/python/squares_of_a_sorted_array.py) | [✅](./array/ts/squares_of_a_sorted_array.ts) |
| 1200 | 最小绝对差 | [✅](./array/golang/minimum_absolute_difference.go) | [✅](./array/python/minimum_absolute_difference.py) | [✅](./array/ts/minimum_absolute_difference.ts) |
| 1287 | 有序数组中出现次数超过 25% 的元素 | [✅](./array/golang/element_appearing_more_than-_5-_n_sorted_array.go) | [✅](./array/python/element_appearing_ore_than_25_in_sorted_array.py) | [✅](./array/ts/element_appearing_ore_than_25_in_sorted_array.ts) |

> 💡 **88** 用「从后往前」归并，避免覆盖 `nums1` 尚未处理的元素；**349** 先排序再双指针并**跳过重复值**，保证结果唯一；**977** 平方后最大值只可能来自首尾两端，双指针从两端往中间、结果从后往前填。
>
> 📎 88 另附 Go 单元测试：[测试](./array/golang/merge_sorted_array_test.go)

### 综合 · 分区、翻转与轮转

| 题号 | 题目 | Go | Python | TS |
| ---- | ---- | :-: | :-: | :-: |
| 75 | 颜色分类 | [✅](./array/golang/sort_color.go) | [✅](./array/python/sort_color.py) | [✅](./array/ts/sort_color.ts) |
| 189 | 轮转数组 | [✅](./array/golang/rotate_array.go) | [✅](./array/python/rotae_array.py) | [✅](./array/ts/rotate_array.ts) |

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
