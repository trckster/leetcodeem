#include <iostream>
#include <string>
#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    int findMaxOddPalindrome(string s, int center, int i) {
        for (; center + i < s.length() && center - i >= 0; ++i) {
            if (s[center + i] != s[center - i]) {
                break;
            }
        }

        return i;
    }

    int findMaxEvenPalindrome(string s, int beforeCenter, int i) {
        for (; beforeCenter + 1 + i < s.length() && beforeCenter - i >= 0; ++i) {
            if (s[beforeCenter + 1 + i] != s[beforeCenter - i]) {
                break;
            }
        }

        return i;
    }

    string getLongestOddPalindrome(string s) {
        int l = 0, r = -1, reversedInScope;
        vector<int> d(s.length(), 0);

        for (int i = 0; i < s.length(); ++i) {
            if (i > r) {
                d[i] = this->findMaxOddPalindrome(s, i, 1);
            } else {
                reversedInScope = r - i + l;

                if (d[reversedInScope] > reversedInScope - l) {
                    d[i] = r - i;
                    d[i] = this->findMaxOddPalindrome(s, i, d[i]);
                } else {
                    d[i] = d[reversedInScope];
                }
            }

            if (r < i + d[i] - 1) {
                l = i - d[i] + 1;
                r = i + d[i] - 1;
            }
        }

        int position = 0;
        int len = 1;
        for (int i = 0; i < s.length(); ++i) {
            if (d[i] > len) {
                position = i;
                len = d[i];
            }
        }

        return s.substr(position - len + 1, len * 2 - 1);
    }

    string getLongestEvenPalindrome(string s) {
        int l = -1, r = -1, reversedInScope;
        int vectorSize = s.length() - 1;
        vector<int> d(vectorSize, 0);

        for (int i = 0; i < vectorSize; ++i) {
            if (i > r) {
                d[i] = this->findMaxEvenPalindrome(s, i, 0);
            } else {
                reversedInScope = r - i + l;

                if (d[reversedInScope] > r - i) {
                    d[i] = r - i;
                    d[i] = this->findMaxEvenPalindrome(s, i, d[i]);
                } else {
                    d[i] = d[reversedInScope];
                }
            }

            if (d[i] && r < i + d[i] - 1) {
                l = i - d[i] + 1;
                r = i + d[i] - 1;
            }
        }

        int position = 0;
        int len = 0;
        for (int i = 0; i < vectorSize; ++i) {
            cout << d[i] << "\n";
            if (d[i] > len) {
                position = i;
                len = d[i];
            }
        }

        return s.substr(position - len + 1, len * 2);
    }

    string longestPalindrome(string s) {
        string odd = getLongestOddPalindrome(s);
        string even = getLongestEvenPalindrome(s);

        if (odd.length() > even.length()) {
            return odd;
        }

        return even;
    }
};

int main() {
    Solution solution;

    string toSolve;

    cin >> toSolve;
    cout << solution.longestPalindrome(toSolve) << "\n";
}