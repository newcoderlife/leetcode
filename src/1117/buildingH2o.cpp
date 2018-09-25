#include <functional>
#include <mutex>
#include <thread>

using namespace std;

class H2O {
 private:
  mutex h_mtx, o_mtx;
  int counter;

 public:
  H2O() {
    o_mtx.lock();
    counter = 2;
  }

  void hydrogen(function<void()> releaseHydrogen) {
    h_mtx.lock();
    releaseHydrogen();
    counter--;
    if (counter == 0)
      o_mtx.unlock();
    else
      h_mtx.unlock();
  }

  void oxygen(function<void()> releaseOxygen) {
    o_mtx.lock();
    releaseOxygen();
    counter = 2;
    h_mtx.unlock();
  }
};

int main(int argc, char const *argv[]) { return 0; }
