function targetIndices(nums: number[], target: number): number[] {
      let less =  0;
      let equal = 0;
      
      for(const item of nums) {
        if (item < target) {
            less +=1;
        } else if (item == target ){
            equal +=1;
        }
      }

    const ans: number[] = new Array(equal)

    for (let i = 0;i < equal;i++) {
        ans[i] = i + less;
    }

    return ans;
};