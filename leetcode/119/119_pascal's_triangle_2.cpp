#include <iostream>
#include <vector>
using namespace std;

class Solution
{
public:
    vector<int> getRow(int rowIndex)
    {
        vector<int> v(rowIndex + 1);
        v.at(0) = 1;
        if (rowIndex == 0)
        {
            return v;
        }
        v.at(1) = 1;
        if (rowIndex == 1)
        {
            return v;
        }

        for (int i = 2; i < rowIndex + 1; i++)
        {
            for (int j = i - 1; j > 0; j--)
            {
                v.at(j) = v.at(j - 1) + v.at(j);
            }
            v.at(i) = 1;
        }
        return v;
    }
};

int main()
{
    return 0;
}