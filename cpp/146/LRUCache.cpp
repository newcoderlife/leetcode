#include <iostream>
#include <list>
#include <unordered_map>

using namespace std;

class LRUCache {
 private:
  int cap;
  unordered_map<int, list<pair<int, int>>::iterator> exists;
  list<pair<int, int>> cache;

 public:
  LRUCache(int capacity) : cap(capacity) {}

  int get(int key) {
    auto it = exists.find(key);
    if (it == exists.end()) return -1;

    pair<int, int> node = *exists[key];
    cache.erase(exists[key]);

    cache.push_front(node);
    exists[key] = cache.begin();

    return node.second;
  }

  void put(int key, int value) {
    auto it = exists.find(key);
    if (it == exists.end()) {
      if (cache.size() == this->cap) {
        exists.erase(cache.back().first);
        cache.pop_back();
      }
    } else
      cache.erase(exists[key]);

    cache.push_front(make_pair(key, value));
    exists[key] = cache.begin();
  }
};

int main(int argc, char *argv[]) {
  LRUCache cache(3);
  for (int i = 0; i < 10; i++) {
    cache.put(i, i);
  }
  for (int i = 0; i < 10; i++) {
    cout << cache.get(i) << endl;
  }
  return 0;
}