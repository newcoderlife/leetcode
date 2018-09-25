#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  int countNegatives(vector<vector<int>>& grid) {
    if (grid.empty()) return 0;
    int n = grid.size(), m = grid.front().size(), count = 0;

    for (int i = 0, j; i < n; i++, count += m - j)
      for (j = 0; j < m; j++)
        if (grid[i][j] < 0) break;

    return count;
  }
};

int main(int argc, char const* argv[]) {
  cout << Solution().countNegatives(vector<vector<int>>{
              {4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}})
       << endl;
  return 0;
}
