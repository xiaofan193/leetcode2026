function nthUglyNumber(n: number): number {
    const ans: number[] = []
    ans[1] =1
    let p2 = 1
    let p3 =1
    let p5 = 1

    for (let i=2;i < n+1;i++){
        let  a = ans[p2]* 2
        let b = ans[p3]*3
        let c = ans[p5]*5

       let  mind =Math.min(a,Math.min(b,c))

       if (mind === a) {
            p2+=1
       }

       if (mind===b) {
            p3 +=1
       }
       
       if (mind ===c) {
            p5 +=1
       }

       ans[i] = mind 

    }

    return ans[n]
};