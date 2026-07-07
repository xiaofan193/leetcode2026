package backtracking

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

var (
    m []string
    path []byte
    res []string
)
func letterCombinations(digits string) []string {
    m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    path, res = make([]byte, 0), make([]string, 0)
    if digits == "" {
        return res
    }
    dfs(digits, 0)
    return res
}
func dfs(digits string, start int) {
    if len(path) == len(digits) {  // Termination condition, string length is equal to the length of digits
        tmp := string(path)
        res = append(res, tmp)
        return
    }
    digit := int(digits[start] - '0')  // Convert the digit pointed by index into an int (confirm the next digit)
    str := m[digit-2]   // Get the character set corresponding to the digit (note the mapping)
    for j := 0; j < len(str); j++ {
        path = append(path, str[j])
        dfs(digits, start+1)
        path = path[:len(path)-1]
    }
}