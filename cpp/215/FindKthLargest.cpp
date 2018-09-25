#include <iostream>
#include <vector>
#include <algorithm>
#include <queue>

using namespace std;

class Solution
{
public:
    int findKthLargest(vector<int> &nums, int k)
    {
        vector<int> left, right, mid;
        for (auto &num : nums)
        {
            if (num > nums[0])
                left.push_back(num);
            else if (num == nums[0])
                mid.push_back(num);
            else
                right.push_back(num);
        }

        if (k <= left.size())
            return findKthLargest(left, k);
        else if (k <= left.size() + mid.size())
            return mid[0];
        else
            return findKthLargest(right, k - left.size() - mid.size());
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    cout << sol.findKthLargest(vector<int>{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) << endl;
    cout << sol.findKthLargest(vector<int>{3, 2, 1, 5, 6, 4}, 2) << endl;
    return 0;
}