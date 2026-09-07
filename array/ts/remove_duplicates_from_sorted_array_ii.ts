function removeDuplicates2(nums: number[]): number {
    const n = nums.length;
    if (n === 1) return 1;
    let k = 2;
    for (let i = 2; i < n; i++) {
        if (!((nums[k - 2] === nums[k - 1]) && (nums[i] === nums[k - 1]))) {
            nums[k] = nums[i];
            k++;
        }
    }
    return k;
}
