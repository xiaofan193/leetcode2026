package backtracing

var (
	path []int
	res [][]int
)
func combine(n int, i int) [][]int {
	 res, path = make([][]int, 0), make([]int, 0, k)
     backtracing(k, n, 1, 0)
    return res
}

func backtracing(n int ,k int ,start int,sum int ) {
	if len(path) == k {
		tmp := make([]int,k)
		copy(tmp,path)
		res = append(res,tmp)
		return  
	}

	for i := start; i <= 9; i++ {
		if sum + i > n || 9 -i + 1 < k -len(path) {
			break 
		}
		path  = append(path,i)
		backtracing(n,k,i+1,sum+i)
		path = path[: len(path)-1]
	}
}