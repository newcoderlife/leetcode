#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>

using namespace std;

class Solution
{
public:
    int minimumTotal(vector<vector<int>> &triangle)
    {
        if (triangle.size() == 0)
            return 0;

        for (int i = triangle.size() - 2; i >= 0; i--)
            for (int j = 0; j < triangle[i].size(); j++)
                triangle[i][j] = min(triangle[i + 1][j], triangle[i + 1][j + 1]) + triangle[i][j];

        return triangle[0][0];
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    cout << sol.minimumTotal(vector<vector<int>>{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}) << endl;
    return 0;
}