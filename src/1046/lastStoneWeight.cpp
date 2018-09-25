#include <algorithm>
#include <iostream>
#include <queue>
#include <vector>

using namespace std;

class Solution {
 public:
  int lastStoneWeight(vector<int>& stones) {
    priority_queue<int, vector<int>, less<int>> heap;
    for (int i : stones) {
      heap.push(i);
    }

    while (heap.size() > 1) {
      int a = heap.top();
      heap.pop();
      int b = heap.top();
      heap.pop();

      if (a != b) {
        heap.push(abs(a - b));
      }
    }
    if (heap.empty()) {
      return 0;
    }
    return heap.top();
  }
};

int main(int argc, char const* argv[]) { return 0; }
