
// 93. 复原 IP 地址
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

 

// 示例 1：

// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 示例 2：

// 输入：s = "0000"
// 输出：["0.0.0.0"]
// 示例 3：

// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

// https://leetcode.cn/problems/restore-ip-addresses/description

var (
    path []string
    res  []string
)
func restoreIpAddresses(s string) []string {
    path, res = make([]string, 0, len(s)), make([]string, 0)
    dfs(s, 0)
    return res
}
func dfs(s string, start int) {  
    if len(path) == 4 {    // Don't continue if there are already four segments
        if start == len(s) {      
            str := strings.Join(path, ".")
            res = append(res, str)
        }
        return 
    }
    for i := start; i < len(s); i++ {
        if i != start && s[start] == '0' { // Segments with leading zeros are invalid
            break
        }
        str := s[start : i+1]
        num, _ := strconv.Atoi(str)
        if num >= 0 && num <= 255 {
            path = append(path, str)  // Move to the next layer if valid
            dfs(s, i+1)
            path = path[:len(path) - 1]
        } else {   // If not valid, move directly to exit
            break
        }
    }
}