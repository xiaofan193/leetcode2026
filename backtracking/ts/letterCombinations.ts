// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
// 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

const letterMap: string[] = ["abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"];

function letterCombinations(digits: string): string[] {
    const resArr: string[] = [];
    if (digits === "") return resArr;

    backTracking(digits, 0, []);
    return resArr;

    function backTracking(digits: string, start: number, route: string[]): void {
        // 终止条件：组合长度等于 digits 长度
        if (route.length === digits.length) {
            resArr.push(route.join(""));
            return;
        }
        // 获取当前数字对应的字母集合（注意映射：digit 2 -> index 0）
        const str: string = letterMap[Number(digits[start]) - 2];
        for (let i = 0; i < str.length; i++) {
            route.push(str[i]);
            backTracking(digits, start + 1, route);
            route.pop();
        }
    }
}
