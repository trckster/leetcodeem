#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    void setSeparators(string &s) {
        int openedBracketsCount = 0;

        for (int i = 0; i < s.size(); ++i) {
            if (s[i] == '(') {
                openedBracketsCount++;
            } else if (openedBracketsCount > 0) {
                openedBracketsCount--;
            } else {
                s[i] = 's';
            }
        }

        openedBracketsCount = 0;

        for (int i = s.size() - 1; i >= 0; --i) {
            if (s[i] == 's') {
                continue;
            }

            if (s[i] == ')') {
                openedBracketsCount++;
            } else if (openedBracketsCount > 0) {
                openedBracketsCount--;
            } else {
                s[i] = 's';
            }
        }
    }

    int longestValidParentheses(string s) {
        int answer = 0, localAnswer = 0;

        this->setSeparators(s);

        for (int i = 0; i < s.size(); ++i) {
            if (s[i] == 's') {
                localAnswer = 0;
            } else {
                localAnswer++;
            }

            answer = max(localAnswer, answer);
        }

        return answer;
    }
};

int main() {
    Solution s;
    cout << s.longestValidParentheses("") << endl;
    cout << s.longestValidParentheses("(") << endl;
    cout << s.longestValidParentheses(")") << endl;
    cout << s.longestValidParentheses("()") << endl;
    cout << s.longestValidParentheses(")(") << endl;
    cout << s.longestValidParentheses("(()") << endl;
    cout << s.longestValidParentheses("(()())") << endl;
    cout << s.longestValidParentheses("(()()()") << endl;
    cout << s.longestValidParentheses(")()())") << endl;
    cout << s.longestValidParentheses("()(()") << endl;
}