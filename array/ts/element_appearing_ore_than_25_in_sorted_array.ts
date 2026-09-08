function findSpecialInteger(arr: number[]): number {
    const n: number= arr.length;
    let c = arr[0];
    let k = 0;

    for (let i = 0;i < n;i++) {
        if (arr[i] === c) {
            k+=1
            if (k*4 > n ){
                return c
            }
        }else{
            c = arr[i]
            k = 1
        }
    }
    return c
}