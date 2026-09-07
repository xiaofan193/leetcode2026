
// 283
function moveZeroes(nums: number[]): void {
    let n: number = nums.length;
    let k: number = 0
    
    for (let i =0;i < n;i++) {
        if (nums[i] != 0) {
            nums[k] = nums[i];
            k++
        }
    }

    while(k < n){
        nums[k] = 0;
        k++
    }
    

}