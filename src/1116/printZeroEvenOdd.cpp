#include <atomic>
#include <functional>
#include <mutex>
#include <thread>

using namespace std;

class ZeroEvenOdd {
 private:
  int n;
  bool next_odd;
  mutex mtx_odd, mtx_even, mtx_zero;

 public:
  ZeroEvenOdd(int n) {
    this->n = n;
    next_odd = true;
    mtx_odd.lock();
    mtx_even.lock();
  }

  void zero(function<void(int)> printNumber) {
    for (int i = 0; i < n; i++) {
      mtx_zero.lock();
      printNumber(0);
      if (next_odd)
        mtx_odd.unlock();
      else
        mtx_even.unlock();
    }
  }

  void even(function<void(int)> printNumber) {
    for (int i = 2; i <= n; i += 2) {
      mtx_even.lock();
      printNumber(i);
      next_odd = true;
      mtx_zero.unlock();
    }
  }

  void odd(function<void(int)> printNumber) {
    for (int i = 1; i <= n; i += 2) {
      mtx_odd.lock();
      printNumber(i);
      next_odd = false;
      mtx_zero.unlock();
    }
  }
};

int main(int argc, char const *argv[]) { return 0; }
