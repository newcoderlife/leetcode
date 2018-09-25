#include <iostream>
#include <vector>
#include <numeric>

using namespace std;

class Solution
{
public:
    void merge(vector<int> &nums1, int m, vector<int> &nums2, int n)
    {
        vector<int> num;
        num.reserve(m);

        for (int i = 0, j = 0; i < m || j < n;)
        {
            int n1 = (i < m) ? nums1[i] : numeric_limits<int>::max();
            int n2 = (j < n) ? nums2[j] : numeric_limits<int>::max();
            if (n1 < n2)
            {
                num.push_back(n1);
                i++;
            }
            else
            {
                num.push_back(n2);
                j++;
            }
        }

        nums1.swap(num);
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    vector<int> nums1{1, 2, 3, 0, 0, 0};
    sol.merge(nums1, 3, vector<int>{2, 5, 6}, 2);
    return 0;
}