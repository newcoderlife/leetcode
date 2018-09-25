#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int maxArea(vector<int> &height) {
    int result = 0;
    for (int i = 0, j = height.size() - 1; i < j;) {
      result = max((j - i) * min(height[i], height[j]), result);
      if (height[i] < height[j])
        i++;
      else
        j--;
    }
    return result;
  }
};

int main(int argc, char *argv[]) {
  Solution sol;
  vector<int> height{1, 8, 6, 2, 5, 4, 8, 3, 7};
  cout << sol.maxArea(height) << endl;
  return 0;
}