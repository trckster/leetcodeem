#include <bits/stdc++.h>

using namespace std;

void printVectorOfVectors(vector <vector<int>> &v) {
    for (int i = 0; i < v.size(); ++i) {
        for (int j = 0; j < v[i].size(); ++j) {
            cout << v[i][j] << ' ';
        }
        cout << '\n';
    }
    cout << '\n';
}

void printVector(vector<int> &v) {
    for (int i = 0; i < v.size(); ++i) {
        cout << v[i] << ' ';
    }
    cout << '\n';
}

class Solution {
public:
    map<int, map<int, int>> dp;
    
    int minCost(int n, vector<int> &cuts) {
        cuts.push_back(0);
        cuts.push_back(n);
        sort(cuts.begin(), cuts.end());

        for (int i = 0, from, to; i + 1 < cuts.size(); ++i) {
            from = cuts[i];
            to = cuts[i + 1];
            this->dp[from][to] = 0;
        }

        int currentCut, bestCut;
        for (int blockLength = 2; blockLength <= cuts.size(); ++blockLength) {
            vector<int>::iterator from = cuts.begin();
            vector<int>::iterator to = cuts.begin() + blockLength;

            while (to != cuts.end()) {
                bestCut = INT_MAX;
                for (auto it = from + 1; it != to; it++) {
                    currentCut = dp[*from][*it] + dp[*it][*to] + *to - *from;
                    bestCut = min(bestCut, currentCut);
                }

                this->dp[*from][*to] = bestCut;

                from++;
                to++;
            }
        }

        return this->dp[0][n];
    }
};

int main() {
    int n_1 = 7;
    vector<int> cuts_1{1, 3, 4, 5};

    int n_2 = 9;
    vector<int> cuts_2{5, 6, 1, 4, 2};

    int n_3 = 655;
    vector<int> cuts_3{361, 132, 484, 271, 509, 401, 5, 181, 420, 6, 254, 13, 10, 134, 610, 108, 275, 593, 277, 394,
                       101, 267, 333, 248, 232, 474, 56, 304, 523, 543, 295, 143, 171, 299, 130, 512, 448, 180, 629,
                       294, 502, 384, 82, 35, 164, 496, 620, 40, 637, 352, 513, 374, 450, 25, 195, 43, 276, 398, 414,
                       20, 337, 38, 331, 380, 559, 278, 241, 291, 468, 422, 454, 7, 67, 596, 259, 614, 27, 636, 647,
                       208, 272, 106, 541, 377, 530, 451};

    Solution solution;

    int answer_1 = solution.minCost(n_1, cuts_1),
            answer_2 = solution.minCost(n_2, cuts_2),
            answer_3 = solution.minCost(n_3, cuts_3);

    cout << answer_1 << ' ' << answer_2 << ' ' << answer_3 << '\n';
}
