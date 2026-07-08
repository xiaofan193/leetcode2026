function permute(nums: number[]): number[][] {
    const resArr: number[][] = [];
    const helperSet: Set<number> = new Set();
    backTracking(nums, []);
    return resArr;
    function backTracking(nums: number[], route: number[]): void {
        if (route.length === nums.length) {
            resArr.push([...route]);
            return;
        }
        let tempVal: number;
        for (let i = 0, length = nums.length; i < length; i++) {
            tempVal = nums[i];
            if (!helperSet.has(tempVal)) {
                route.push(tempVal);
                helperSet.add(tempVal);
                backTracking(nums, route);
                route.pop();
                helperSet.delete(tempVal);
            }
        }
    }
};