#include <vector>
#include <map>
#include <iostream>

using namespace std;

class Solution
{
public:
    vector<int> twoSum(vector<int> &nums, int target)
    {
        map<int, int> hash;
        for (int i = 0, j; i < nums.size(); i++)
        {
            int tmp = target - nums[i];
            if (hash.count(tmp) > 0 && i != (j = hash[tmp]))
                return vector<int>{i, j};
            else
                hash[nums[i]] = i;
        }
        return vector<int>();
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    auto result = sol.twoSum(vector<int>{3, 2, 4}, 6);
    if (result.size() == 2)
        cout << result[0] << ' ' << result[1] << endl;
    return 0;
}