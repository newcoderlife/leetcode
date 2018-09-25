#include <functional>
#include <thread>

using namespace std;

class FooBar {
 private:
  int n;
  mutex mtx_1, mtx_2;

 public:
  FooBar(int n) {
    this->n = n;
    mtx_2.lock();
  }

  void foo(function<void()> printFoo) {
    for (int i = 0; i < n; i++) {
      mtx_1.lock();
      printFoo();
      mtx_2.unlock();
    }
  }

  void bar(function<void()> printBar) {
    for (int i = 0; i < n; i++) {
      mtx_2.lock();
      printBar();
      mtx_1.unlock();
    }
  }
};
