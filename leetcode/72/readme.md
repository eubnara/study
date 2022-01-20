
```
dp := [501][501]int{}
```

Runtime: 110 ms
Memory: 4.9 MB

```
al := len(word1)
bl := len(word2)

dp := make([][]int, al)
for i := range dp {
    dp[i] = make([]int, bl)
}
```

Runtime: 14 ms
Memory: 5.9 MB


golang 에선 0으로 (zero value) 초기화하는 비용이 있는지 array 로 미리 큰 배열을 선언하지 않고 make 를 이용해서 필요한 만큼만 확보하는 것이 속도 개선에 영향이 있었다.

https://go.dev/ref/spec#The_zero_value
