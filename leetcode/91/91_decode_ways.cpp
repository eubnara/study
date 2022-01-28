#include <iostream>

using namespace std;

class Solution {
public:
    int numDecodings(string s) {
        int l = s.length();
        int *dp = new int[l]{0};
        int ans = 0;
        if (s[0] == '0')
        {
            ans = dp[0];
            delete[] dp;
            return ans;
        }
        else
        {
            dp[0] = 1;
        }
        if (l == 1) {
            ans = dp[0];
            delete[] dp;
            return ans;
        }
        if (s[1] != '0') {
            dp[1] += 1;
        }
        if (s[0] - '0' == 1) {
            dp[1] += 1;
        }
        if (s[0] - '0' == 2 && s[1] - '0' <= 6) {
            dp[1] += 1;
        }
        int p = dp[0], q = dp[1];
        for (int i = 2; i < l; i++) {
            if (p == 0 && q == 0) {
                delete[] dp;
                return 0;
            }
            int t = s[i] - '0', t2 = s[i - 1] - '0';
            if (t != 0) {
                dp[i] += q;
            }
            if (t2 == 1) {
                dp[i] += p;
            }
            if (t2 == 2 && t <= 6) {
                dp[i] += p;
            }
            p = q;
            q = dp[i];
        }
        ans = dp[l - 1];
        delete[] dp;
        return ans;
    }
};

int main() {
}