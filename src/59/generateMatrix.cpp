#include <iostream>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> generateMatrix(int n) {
    /*
     * 0 right, 1 down, 2 left, 3 up
     */
    const int next_directions[4] = {1, 2, 3, 0};
    const int next_row[4] = {0, 1, 0, -1};
    const int next_col[4] = {1, 0, -1, 0};
    int target = pow(n, 2);
    int current_direction = 0;
    int row = 0, col = -1;
    vector<vector<int>> results;
    for (int i = 0; i < n; i++) results.emplace_back(n, 0);

    for (int i = 0; i < target; i++) {
      row += next_row[current_direction];
      col += next_col[current_direction];
      if (row == -1 || col == -1 || row == n || col == n ||
          results[row][col] != 0) {
        row -= next_row[current_direction];
        col -= next_col[current_direction];
        current_direction = next_directions[current_direction];
        row += next_row[current_direction];
        col += next_col[current_direction];
      }

      results[row][col] = i + 1;
    }

    return results;
  }
};

int main(int argc, char const *argv[]) {
  Solution().generateMatrix(3);
  return 0;
}
