#include <iostream>
#include <vector>
#include <numeric>

using namespace std;

class Solution
{
public:
    vector<vector<int>> subsets(vector<int> &nums)
    {
        int total = 1 << nums.size();
        vector<vector<int>> result(total);
        for (int i = 0; i < total; i++)
            for (int j = i, k = 0; j; j >>= 1, k++)
                if (j & 1 == 1)
                    result[i].push_back(nums[k]);
        return result;
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    auto result = sol.subsets(vector<int>{1, 2, 3});
    return 0;
}