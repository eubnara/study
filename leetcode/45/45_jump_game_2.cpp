#include <iostream>
#include <vector>

using namespace std;

class Solution
{
public:
    int jump(vector<int> &nums)
    {
        int l = nums.size();
        if (l <= 1)
        {
            return 0;
        }
        int maxIdx = 0;
        int step = 0;
        for (int i = 0; i < l - 1; i++)
        {
            int curMaxIdx = maxIdx;
            for (int j = i; j <= maxIdx; j++)
            {
                int v = j + nums[j];
                if (curMaxIdx < v)
                {
                    curMaxIdx = v;
                }
            }
            i = maxIdx;
            step++;
            maxIdx = curMaxIdx;
            if (maxIdx >= l - 1)
            {
                return step;
            }
        }
        return step;
    }
};

int main()
{
    Solution s;
    vector<int> v{2, 3, 1, 1, 4};
    // vector<int> v{1, 2, 3};
    cout << s.jump(v) << endl;
}