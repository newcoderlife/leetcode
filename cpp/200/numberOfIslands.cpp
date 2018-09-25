#include <iostream>
#include <vector>

using namespace std;

class Solution {
 private:
  vector<int> father;
  vector<int> rank;
  int count;

  int find(int x) { return father[x] == x ? x : father[x] = find(father[x]); }

  void unionSet(int x, int y) {
    int x_fa = find(x), y_fa = find(y);
    if (x_fa == y_fa) return;

    if (rank[x_fa] < rank[y_fa])
      father[x_fa] = y_fa;
    else {
      father[y_fa] = x_fa;
      if (rank[x_fa] == rank[y_fa]) rank[x_fa]++;
    }

    count--;
  }

 public:
  int numIslands(vector<vector<char>>& grid) {
    if (grid.size() == 0) return 0;

    int n = grid.size(), m = grid.front().size(), total = n * m;
    count = 0;
    father.resize(n * m, 0);
    rank.resize(n * m, 0);
    for (int i = 0; i < n; i++)
      for (int j = 0; j < m; j++)
        if (grid[i][j] == '1') {
          father[i * m + j] = i * m + j;
          count++;
        } else
          father[i * m + j] = -1;

    for (int i = 0; i < n; i++)
      for (int j = 0; j < m; j++)
        if (grid[i][j] == '1') {
          grid[i][j] = '0';
          if (i > 0 && grid[i - 1][j] == '1')
            unionSet(i * m + j, (i - 1) * m + j);
          if (j > 0 && grid[i][j - 1] == '1')
            unionSet(i * m + j, i * m + j - 1);
          if (i + 1 < n && grid[i + 1][j] == '1')
            unionSet(i * m + j, (i + 1) * m + j);
          if (j + 1 < m && grid[i][j + 1] == '1')
            unionSet(i * m + j, i * m + j + 1);
        }
    return count;
  }
};

int main(int argc, char const* argv[]) {
  Solution sol;
  vector<vector<char>> grid{{'1', '1', '1', '1', '0'},
                            {'1', '1', '0', '1', '0'},
                            {'1', '1', '0', '0', '0'},
                            {'0', '0', '0', '0', '0'}};

  cout << sol.numIslands(grid) << endl;
  return 0;
}
