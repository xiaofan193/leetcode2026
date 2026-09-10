function shuffle(nums: number[], n: number): number[] {
    const ans: number[] = []
   let  i = 0;
   let j = n;

   while(i < n) {
    ans.push(nums[i]);
    ans.push(nums[j]);
        i+=1;
        j+=1;

   }

   return ans;
   
};