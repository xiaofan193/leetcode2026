
// https://leetcode.cn/problems/subsets/description/
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：
// 输入：nums = [0]
// 输出：[[],[0]]
 

// 提示：
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同

var (
	path []int 
	res [][]int
)

func subsets(nums []int) [][]int {
    path = make([]int,0)
	res = make([][]int,0)

	return res 
}

func dfs(nums []int,start int){

	if start == len(nums) {
		return 
	}

	tmp := make([]int,0)
	copy(tmp,path)
	res = append(res,tmp)
		
	
	for i := start;i < len(nums);i++ {
		path = append(path,nums[i])
		dfs(nums,i+1)
		path = path[:len(path) -1]
	}
}


func subsetsWithDup(nums []int) [][]int {
    path, res = make([]int, 0, len(nums)), make([][]int, 0)
    sort.Ints(nums)
    dfs(nums, 0)
    return res
}

func dfs2(nums []int, start int) {
    tmp := make([]int, len(path))
    copy(tmp, path)
    res = append(res, tmp)

    for i := start; i < len(nums); i++ {
        if i != start && nums[i] == nums[i-1] {
            continue
        }
        path = append(path, nums[i])
        dfs(nums, i+1)
        path = path[:len(path)-1]
    }
}