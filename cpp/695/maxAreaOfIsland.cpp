#include <iostream>
#include <vector>

using namespace std;

class Solution {
 private:
  vector<int> father;
  vector<int> rank;
  int result;
  int find(int x) { return father[x] == x ? x : father[x] = find(father[x]); }

  void unionSet(int x, int y) {
    int fa_x = find(x), fa_y = find(y);
    if (fa_x == fa_y) return;

    if (rank[fa_x] < rank[fa_y])
      father[fa_x] = fa_y;
    else
      father[fa_y] = fa_x;

    rank[fa_x] += rank[fa_y];
    rank[fa_y] = rank[fa_x];
    result = max(result, rank[fa_x]);
  }

 public:
  int maxAreaOfIsland(vector<vector<int>>& grid) {
    int n = grid.size();
    if (n == 0) return 0;
    int m = grid.front().size();
    rank.resize(m * n, 0);
    father.resize(n * m, -1);
    result = 0;
    for (int i = 0; i < n; i++)
      for (int j = 0; j < m; j++)
        if (grid[i][j] == 1) {
          father[i * m + j] = i * m + j;
          rank[i * m + j] = 1;
          result = 1;
        }

    for (int i = 0; i < n; i++) {
      for (int j = 0; j < m; j++) {
        if (grid[i][j] == 1) {
          grid[i][j] = 0;

          if (i > 0 && grid[i - 1][j] == 1)
            unionSet(i * m + j, (i - 1) * m + j);

          if (j > 0 && grid[i][j - 1] == 1) unionSet(i * m + j, i * m + j - 1);

          if (i + 1 < n && grid[i + 1][j] == 1)
            unionSet(i * m + j, (i + 1) * m + j);

          if (j + 1 < m && grid[i][j + 1] == 1)
            unionSet(i * m + j, i * m + j + 1);
        }
      }
    }

    return result;
  }
};

int main(int argc, char const* argv[]) {
  vector<vector<int>> grid{{0}};
  cout << Solution().maxAreaOfIsland(grid) << endl;
  return 0;
}
