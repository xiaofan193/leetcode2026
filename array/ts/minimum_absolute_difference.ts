function minimumAbsDifference(arr: number[]): number[][] {
    const res: number[][] = []
    if (arr.length === 0) {
        return res
    }

    // 数字必须传比较器，默认 sort 是按字典序排序
    arr.sort((a, b) => a - b)
    let min = arr[1] - arr[0]

    for (let i = 1; i < arr.length; i++) {
        const item = [arr[i - 1], arr[i]]

        if (arr[i] - arr[i - 1] === min) {
            res.push(item)
        }

        if (arr[i] - arr[i - 1] < min) {
            min = arr[i] - arr[i - 1]
            res.length = 0 // 相当于 Go 的 res[:0]
            res.push(item)
        }
    }
    return res
}
