/**
 Do not return anything, modify nums in-place instead.
 */
function sortColors(nums: number[]): void {
    const n = nums.length;
    const count = [0, 0, 0]; // 统计 0/1/2 各出现几次

    for (let i = 0; i < n; i++) {
        count[nums[i]]++;
    }

    let k = 0;
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < count[i]; j++) {
            nums[k] = i;
            k++;
        }
    }
}
