#include <algorithm>
#include <iostream>
#include <list>
#include <vector>

using namespace std;

class ProductOfNumbers {
 private:
  vector<list<int>> facs;
  int count;

 public:
  ProductOfNumbers() {
    facs.resize(101, list<int>());
    count = 0;
  }

  void add(int num) { facs[num].push_back(count++); }

  int getProduct(int k) {
    int result = 1, left = count - k;
    for (int i = 0, j; i < 101; i++) {
      j = 0;
      for (auto iter = facs[i].begin(); iter != facs[i].end(); iter++, j++) {
        if (*iter >= left) {
          result *= pow(i, facs[i].size() - j);
          if (result == 0) return 0;
          break;
        }
      }
    }
    return result;
  }
};

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * ProductOfNumbers* obj = new ProductOfNumbers();
 * obj->add(num);
 * int param_2 = obj->getProduct(k);
 */

int main(int argc, char const *argv[]) {
  auto sol = ProductOfNumbers();
  sol.add(3);
  sol.add(0);
  sol.add(2);
  sol.add(5);
  sol.add(4);
  cout << sol.getProduct(2) << endl;
  cout << sol.getProduct(3) << endl;
  cout << sol.getProduct(4) << endl;
  sol.add(8);
  cout << sol.getProduct(2) << endl;
  return 0;
}
