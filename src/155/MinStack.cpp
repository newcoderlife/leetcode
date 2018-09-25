#include <cstdint>
#include <vector>

using namespace std;

class MinStack {
 private:
  vector<int64_t> buf;
  int64_t minimum;

 public:
  MinStack() { minimum = 0; }

  void push(int x) {
    if (buf.empty()) {
      buf.push_back(0);
      minimum = x;
      return;
    }

    buf.push_back(x - minimum);
    if (x - minimum < 0L) minimum = x;
  }

  void pop() {
    if (buf.back() < 0L) minimum -= buf.back();
    buf.pop_back();
  }

  int top() { return buf.back() < 0L ? minimum : minimum + buf.back(); }

  int getMin() { return minimum; }
};

int main(int argc, char const* argv[]) {
  MinStack* minStack = new MinStack();
  minStack->push(1);
  minStack->push(2);
  minStack->pop();
  minStack->getMin();
  minStack->pop();
  minStack->getMin();
  minStack->top();
  return 0;
}
