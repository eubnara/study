public class Solution {
    public int maxProfit(int[] prices) {
        int l = prices.length;
        int[] dp = new int[l + 1];
        int ans = 0;
        dp[l - 1] = 0;
        dp[l] = 0;
        int rightMax = prices[l - 1];
        /*
         * i 부터 l-1 까지에서 얻을 수 있는 최댓값을 미리 계산
         * O(n^2) 이 아니라 O(n) 으로 풀 수 있도록 오른쪽부터 계산한다.
         * 오른쪽에서부터 알아낸 최댓값 rightMax 와 현재값 cur 을 뺀 값 혹은 i+1 부터 l-1 까지 구한 최댓값 중 큰 것을 dp[i]
         * 에 저장한다.
         */
        for (int i = l - 2; i >= 0; i--) {
            int cur = prices[i];
            dp[i] = Math.max(dp[i + 1], rightMax - cur);
            rightMax = Math.max(rightMax, cur);
        }

        /*
         * 왼쪽부터 계산하고 가장 작은 값을 leftMin 에 저장해나가면서 온다.
         * 현재값 - leftMin + dp[i+1] 이 지금껏 구한 ans 보다 큰 값이 나온다면 갱신한다.
         */
        int leftMin = prices[0];
        for (int i = 0; i < l; i++) {
            int cur = prices[i];
            leftMin = Math.min(cur, leftMin);
            ans = Math.max(ans, dp[i + 1] + Math.max(cur - leftMin, 0));
        }
        return ans;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int[][] testcases = {
                { 3, 3, 5, 0, 0, 3, 1, 4 },
                { 1, 2, 3, 4, 5 },
                { 7, 6, 4, 3, 1 }
        };
        for (int[] testcase : testcases) {
            System.out.println(s.maxProfit(testcase));
        }
    }
}