#include <condition_variable>
#include <functional>
#include <mutex>
#include <thread>

using namespace std;

class Foo {
 private:
  mutex mtx_1, mtx_2, mtx_3;

 public:
  Foo() {
    mtx_2.lock();
    mtx_3.lock();
  }

  void first(function<void()> printFirst) {
    mtx_1.lock();
    printFirst();
    mtx_2.unlock();
  }

  void second(function<void()> printSecond) {
    mtx_2.lock();
    printSecond();
    mtx_3.unlock();
  }

  void third(function<void()> printThird) {
    mtx_3.lock();
    printThird();
    mtx_1.unlock();
  }
};

int main(int argc, char const *argv[]) {
  /* code */
  return 0;
}
