#include <vector>

using namespace std;

class Solution {
 public:
  bool searchMatrix(vector<vector<int>>& matrix, int target) {
    if (matrix.empty()) return false;

    int n = matrix.size(), m = matrix.front().size();
    int col = m - 1, row = 0;
    while (true) {
      if (col < 0 || row >= n) break;
      if (matrix[row][col] == target) return true;
      if (matrix[row][col] > target)
        col--;
      else
        row++;
    }
    return false;
  }
};

int main(int argc, char const* argv[]) { return 0; }
