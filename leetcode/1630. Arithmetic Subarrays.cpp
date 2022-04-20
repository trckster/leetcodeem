#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    vector<bool> checkArithmeticSubarrays(vector<int> &nums, vector<int> &l, vector<int> &r) {
        vector<bool> answer = vector<bool>();

        for (int i = 0, from, to, diff; i < l.size(); ++i) {
            from = l[i];
            to = r[i] + 1;

            if (to - from == 1) {
                answer.push_back(true);
                continue;
            }

            vector<int> tmp(to - from);

            partial_sort_copy(nums.begin() + from, nums.begin() + to, tmp.begin(), tmp.end());

            diff = tmp[1] - tmp[0];
            bool diff_ok = true;
            for (int j = 2; j < tmp.size(); ++j) {
                if (tmp[j] - tmp[j - 1] != diff) {
                    diff_ok = false;
                    break;
                }
            }
            answer.push_back(diff_ok);
        }

        return answer;
    }
};

int main() {
    Solution s;

    vector<int> nums{4, 6, 5, 9, 3, 7};
    vector<int> l{0, 0, 2};
    vector<int> r{2, 3, 5};

    vector<bool> answer = s.checkArithmeticSubarrays(nums, l, r);

    for (auto result: answer) {
        cout << (result ? '+' : '-');
    }
    cout << endl;
}