#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>

using namespace std;

class Solution
{
public:
    vector<int> spiralOrder(vector<vector<int>> &matrix)
    {
        vector<int> result;
        if (matrix.size() == 0)
            return result;

        int m = matrix.size(), n = matrix[0].size(), depth = (min(m, n) + 1) / 2;
        result.reserve(m * n);

        for (int i = 0; i < depth; i++)
        {
            for (int j = i; j < n - i; j++)
                result.push_back(matrix[i][j]);
            for (int j = i + 1; j < m - i; j++)
                result.push_back(matrix[j][n - i - 1]);
            for (int j = n - i - 2; j >= i && (m != i * 2 + 1); j--)
                result.push_back(matrix[m - i - 1][j]);
            for (int j = m - i - 2; j > i && (n != i * 2 + 1); j--)
                result.push_back(matrix[j][i]);
        }

        return result;
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    auto result = sol.spiralOrder(vector<vector<int>>{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}});
    for (auto i : result)
        cout << i << " ";
    cout << endl;
    result = sol.spiralOrder(vector<vector<int>>{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}});
    for (auto i : result)
        cout << i << " ";
    cout << endl;
    return 0;
}