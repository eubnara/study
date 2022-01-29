#include <vector>
#include <iostream>
using namespace std;

class Solution
{
public:
    int numTrees(int n)
    {
        vector<vector<int>> dp(n + 1, vector<int>(n + 1, 0));
        return numTrees(1, n, dp);
    }
    int numTrees(int start, int end, vector<vector<int>> &dp)
    {
        if (start == end)
        {
            return 1;
        }
        if (dp.at(start).at(end) != 0)
        {
            return dp.at(start).at(end);
        }
        int ans = 0;
        for (int i = start; i <= end; i++)
        {
            int left = 1, right = 1;
            if (start + 1 < i)
            {
                left = numTrees(start, i - 1, dp);
            }
            if (i < end - 1)
            {
                right = numTrees(i + 1, end, dp);
            }
            ans += left * right;
        }
        dp.at(start).at(end) = ans;
        return ans;
    }
};

int main()
{
    Solution s;
    cout << s.numTrees(19) << endl;
}