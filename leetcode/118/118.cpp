#include <vector>
#include <iostream>

using namespace std;

class Solution
{
public:
    vector<vector<int>> generate(int numRows)
    {
        vector<vector<int>> ans;
        ans.push_back(vector<int>{1});
        if (numRows >= 2)
        {
            ans.push_back(vector<int>{1, 1});
        }
        for (int i = 3; i <= numRows; i++)
        {
            vector<int> subAns{1};
            for (int j = 1; j < i - 1; j++)
            {
                subAns.push_back(ans[i - 2][j - 1] + ans[i - 2][j]);
            }
            subAns.push_back(1);
            ans.push_back(subAns);
        }
        return ans;
    }
};

int main()
{
    Solution s = Solution();
    for (auto i : s.generate(5))
    {
        for (auto j : i)
        {
            cout << j << " ";
        }
        cout << endl;
    }
}