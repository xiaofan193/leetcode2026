// 905 按奇偶排序数组
// https://leetcode.cn/problems/sort-array-by-parity/
// 把偶数移到前面、奇数移到后面，不要求保持原相对顺序，就地修改即可。

// 双指针：i 从左找"最左边的奇数"，j 从右找"最右边的偶数"，
// 两个都找到就交换（一个奇数换到右边、一个偶数换到左边），然后问题缩小为 [i+1, j-1] 的子问题。
// 注意判断奇数必须用 `% 2 !== 0` 而不是 `=== 1`：JS 的取模对负数返回 -1（-3 % 2 === -1），
// 写成 `=== 1` 会把负奇数误判成偶数。这也是 Go 版 `nums[j]%2 == 1` 的隐患
// （好在题目限制 nums[i] >= 0，才没暴露）。
function sortArrayByParity(nums: number[]): number[] {
    let i = 0;
    let j = nums.length - 1;
    while (i < j) { // 循环直到不足两个数
        if (nums[i] % 2 === 0) { // 当前是偶数，已经在正确一侧
            i++;
        } else if (nums[j] % 2 !== 0) { // 当前是奇数，已经在正确一侧
            j--;
        } else { // nums[i] 是奇数、nums[j] 是偶数，交换
            [nums[i], nums[j]] = [nums[j], nums[i]];
            i++;
            j--;
        }
    }
    return nums;
}
