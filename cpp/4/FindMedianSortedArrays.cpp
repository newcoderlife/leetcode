#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution
{
public:
    double findMedianSortedArrays(vector<int> &nums1, vector<int> &nums2)
    {
        int m = nums1.size(), n = nums2.size();
        if (m > n)
            return findMedianSortedArrays(nums2, nums1);

        int iMin = 0, iMax = m, total = m + n, half = (total + 1) / 2;
        while (iMin <= iMax)
        {
            int i = (iMin + iMax) / 2, j = half - i;

            if (i < iMax && nums2[j - 1] > nums1[i])
                iMin = i + 1;
            else if (i > iMin && nums1[i - 1] > nums2[j])
                iMax = i - 1;
            else
            {
                int maxLeft = 0, minRight = 0;
                if (i == 0)
                    maxLeft = nums2[j - 1];
                else if (j == 0)
                    maxLeft = nums1[i - 1];
                else
                    maxLeft = max(nums1[i - 1], nums2[j - 1]);
                if ((total)&1 == 1)
                    return maxLeft;

                if (i == m)
                    minRight = nums2[j];
                else if (j == n)
                    minRight = nums1[i];
                else
                    minRight = min(nums1[i], nums2[j]);
                return (maxLeft + minRight) / 2.0;
            }
        }

        return 0;
    }
};

int main(int argc, char *argv[])
{
    Solution sol;
    cout << sol.findMedianSortedArrays(vector<int>{}, vector<int>{1}) << endl;
    return 0;
}
