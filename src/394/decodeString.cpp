#include <iostream>
#include <string>
#include <vector>

using namespace std;

string operator*(const string &s, int n) {
  string result;
  for (int i = 0; i < n; i++) {
    result.append(s);
  }
  return result;
}

struct Node {
  int num;
  string str;
  vector<Node *> nexts;

  Node(const int &num, const string &str) : num(num), str(str) {}
};

class Solution {
 public:
  string decodeString(string s) {
    auto root = new Node(1, s);
    buildTree(root);
    return calcNode(root);
  }

 private:
  bool allAlpha(string str) {
    for (auto c : str) {
      if (!isalpha(c)) {
        return false;
      }
    }
    return true;
  }

  string calcNode(Node *cur) {
    if (cur->nexts.empty()) return cur->str * cur->num;

    for (Node *next : cur->nexts) {
      cur->str += calcNode(next);
    }
    return cur->str * cur->num;
  }

  void buildTree(Node *cur) {
    if (allAlpha(cur->str)) return;

    for (int i = 0, j, k; i < cur->str.size(); i++) {
      int next_num = 1;
      string next_str;

      if (isdigit(cur->str[i])) {
        for (j = i + 1; j < cur->str.size() && isdigit(cur->str[j]); j++)
          ;
        next_num = stoi(cur->str.substr(i, j - i));

        int op_cnt = 0;
        for (k = j; k < cur->str.size(); k++) {
          if (cur->str[k] == '[') op_cnt++;
          if (cur->str[k] == ']') op_cnt--;
          if (op_cnt == 0) {
            next_str = cur->str.substr(j + 1, k - j - 1);
            break;
          }
        }

        cur->str.erase(i, k - i + 1);
        i = -1;
      } else if (isalpha(cur->str[i])) {
        for (j = i + 1; j < cur->str.size() && isalpha(cur->str[j]); j++)
          ;
        next_str = cur->str.substr(i, j - i);

        cur->str.erase(i, j - i);
        i = -1;
      }

      cur->nexts.push_back(new Node(next_num, next_str));
    }

    for (Node *next : cur->nexts) {
      buildTree(next);
    }
  }
};

int main(int argc, char const *argv[]) {
  cout << Solution().decodeString("2[b4[F]c]") << endl;
  return 0;
}
