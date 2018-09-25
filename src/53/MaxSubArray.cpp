#include <algorithm>
#include <iostream>
#include <numeric>
#include <vector>

using namespace std;

class Solution {
 public:
  int maxSubArray(const vector<int> &nums) {
    return binarySearch(nums, 0, nums.size());
  }

  int binarySearch(const vector<int> &content, int left, int right) {
    if (left + 1 == right) return content[left];
    int mid = (left + right) / 2;
    int leftMax = binarySearch(content, left, mid);
    int rightMax = binarySearch(content, mid, right);
    int midMax = getMidMax(content, left, right, mid);
    return max(leftMax, max(rightMax, midMax));
  }

  int getMidMax(const vector<int> &content, int left, int right, int mid) {
    int leftMax = numeric_limits<int>::min();
    for (int i = mid, sum = 0; i >= left; i--)
      leftMax = max(leftMax, sum += content[i]);

    int rightMax = numeric_limits<int>::min();
    for (int i = mid, sum = 0; i < right; i++)
      rightMax = max(rightMax, sum += content[i]);

    return leftMax + rightMax - content[mid];
  }
};

int main(int argc, char *argv[]) {
  Solution sol;
  cout << sol.maxSubArray(vector<int>{-2, 1, -3, 4, -1, 2, 1, -5, 4}) << endl;
  cout << sol.maxSubArray(vector<int>{1}) << endl;
  cout << sol.maxSubArray(vector<int>{1, 2}) << endl;
  cout << sol.maxSubArray(vector<int>{-1, -2}) << endl;
  return 0;
}