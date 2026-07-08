function subsets(nums: number[]): number[][] {
    const resArr: number[][] = [];
    backTracking(nums,0,[])
    return resArr;

    function backTracking(nums: number[],start: number,route: number[]): void {
        resArr.push([...route]);
        let lenth = nums.length;
        if (start == lenth) return;
        for(let i= start;i< lenth;i++) {
            route.push(nums[i]);
            backTracking(nums,i+1,route);
            route.pop()
        }
    }
}