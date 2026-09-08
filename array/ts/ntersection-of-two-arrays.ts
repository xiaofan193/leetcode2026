function intersection(nums1: number[], nums2: number[]): number[] {
    
    const res: number[] = [];
    const m: number = nums1.length;
    const n: number = nums2.length;

    let i= 0;
    let j= 0;
    nums1.sort((a,b) => a-b)
    nums2.sort((a,b)=>a-b)
    while (i < m && j <n){
        if (nums1[i]< nums2[j]){
            i++;
        }else if (nums1[i] > nums2[j]) {
            j++ 
        }else {
            if ((res.length ===0) || (res[res.length-1] != nums1[i])){
                res.push(nums1[i]);
                
            }

            i +=1
            j +=1
          

        }
    }
    return res
};