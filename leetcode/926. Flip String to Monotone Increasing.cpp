#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    int minFlipsMonoIncr(string s) {
        int firstOneIndex = -1;
        for (int i = 0; i <= s.length(); i++) {
            if (i == s.length()) {
                return 0;
            }

            if (s[i] == '1') {
                firstOneIndex = i;

                break;
            }
        }

        int lastZeroIndex = -1;
        for (int i = s.length() - 1; i >= -1; i--) {
            if (i == -1) {
                return 0;
            }

            if (s[i] == '0') {
                lastZeroIndex = i;

                break;
            }
        }

        int zeros = 0;

        for (int i = firstOneIndex; i <= lastZeroIndex; i++) {
            if (s[i] == '0') {
                zeros++;
            }
        }

        int answer = INT_MAX, ones = 0;

        for (int splitAfter = firstOneIndex - 1; splitAfter <= lastZeroIndex; splitAfter++) {
            answer = min(answer, ones + zeros);

            if (s[splitAfter + 1] == '0') {
                zeros--;
            } else {
                ones++;
            }
        }

        return answer;
    }
};


int main() {
    Solution s;

    cout << s.minFlipsMonoIncr("001010011") << endl;
    cout << s.minFlipsMonoIncr("1111") << endl;
    cout << s.minFlipsMonoIncr("11110000") << endl;
    cout << s.minFlipsMonoIncr("0000") << endl;
    cout << s.minFlipsMonoIncr("0") << endl;
    cout << s.minFlipsMonoIncr("1") << endl;
    cout << s.minFlipsMonoIncr("01") << endl;
    cout << s.minFlipsMonoIncr("10") << endl;
    cout << s.minFlipsMonoIncr("10011111110010111011") << endl;
}