function climbStairs(n: number): number {
  if (n ===0 ||n ==1 ) {
    return n
  }
  let dp: number[] = [];
  dp[1] =1;
  dp[2] = 2;
  dp[3] =3;
  for(let i = 4;i <=n;i++) {
    dp[i] = dp[i-2] + dp[i-1]
  }
  
  
  return dp[n];
};