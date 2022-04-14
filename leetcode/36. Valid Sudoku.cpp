#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    bool isValidSudoku(vector <vector<char>> &board) {
        set<int> unique_numbers;

        /** Test rows */
        for (auto row: board) {
            for (auto number: row) {
                if (number == '.') {
                    continue;
                }

                if (unique_numbers.find(number) != unique_numbers.end()) {
                    return false;
                }

                unique_numbers.insert(number);
            }

            unique_numbers.clear();
        }

        /** Test columns */
        for (int i = 0; i < board.size(); ++i) {
            for (int j = 0; j < board.size(); ++j) {
                auto number = board[j][i];

                if (number == '.') {
                    continue;
                }

                if (unique_numbers.find(number) != unique_numbers.end()) {
                    return false;
                }

                unique_numbers.insert(number);
            }

            unique_numbers.clear();
        }

        /** Test squares */
        for (int i = 0; i < board.size(); i += 3) {
            for (int j = 0; j < board.size(); j += 3) {
                for (int localI = i; localI < i + 3; ++localI) {
                    for (int localJ = j; localJ < j + 3; ++localJ) {
                        auto number = board[localI][localJ];

                        if (number == '.') {
                            continue;
                        }

                        if (unique_numbers.find(number) != unique_numbers.end()) {
                            return false;
                        }

                        unique_numbers.insert(number);
                    }
                }

                unique_numbers.clear();
            }
        }

        return true;
    }
};

int main() {
    Solution s;

    vector <vector<char>> board1{
            {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
            {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
            {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
            {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
            {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
            {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
            {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
            {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
            {'.', '.', '.', '.', '8', '.', '.', '7', '9'}
    };

    if (s.isValidSudoku(board1)) {
        cout << "Good\n";
    } else {
        cout << "Bad\n";
    }
}