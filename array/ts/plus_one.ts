// 66 加一
// https://leetcode.cn/problems/plus-one/
// 从最低位开始逐位加 1，k 表示当前进位。
// 个位单独处理（9 则写 0 进 1，否则直接 +1），其余位从右向左依次处理：
// 注意必须先把 digits[i] + k 写入结果、再清空 k，否则进位会丢失。
function plusOne(digits: number[]): number[] {
    const n = digits.length;
    const ans: number[] = [];

    let k = 0; // 当前进位

    if (digits[n - 1] === 9) {
        ans.push(0);
        k = 1;
    } else {
        ans.push(digits[n - 1] + 1);
        k = 0;
    }

    for (let i = n - 2; i >= 0; i--) {
        if (digits[i] + k === 10) {
            ans.push(0);
            k = 1;
        } else {
            ans.push(digits[i] + k);
            k = 0;
        }
    }

    if (k === 1) {
        ans.push(1);
    }

    return ans.reverse();
}
