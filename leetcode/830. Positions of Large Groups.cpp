#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    vector <vector<int>> largeGroupPositions(string s) {
        vector <vector<int>> answer;
        s = '-' + s + '-';

        char current_letter = 0;
        int current_best = 0;
        for (int i = 0; i < s.length(); ++i) {
            if (s[i] == current_letter) {
                current_best += 1;
                continue;
            } else if (current_best >= 3) {
                vector<int> range{i - current_best - 1, i - 2};
                answer.push_back(range);
            }

            current_best = 1;
            current_letter = s[i];
        }

        return answer;
    }
};

void printAnswer(vector <vector<int>> v) {
    cout << '[';
    for (auto range: v) {
        cout << "[" << range[0] << ", " << range[1] << "], ";
    }
    cout << "]\n";
}

int main() {
    Solution s;

    printAnswer(s.largeGroupPositions("abbbbaaaaaa"));
    printAnswer(s.largeGroupPositions("aaa"));
}