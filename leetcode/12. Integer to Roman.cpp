#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
    string intToRoman(int num) {
        string answer = "";

        int ones = num % 10;

        if (ones <= 3) {
            answer.append(ones, 'I');
        } else if (ones == 4) {
            answer += "IV";
        } else if (ones < 9) {
            answer += "V";
            answer.append(ones - 5, 'I');
        } else if (ones == 9) {
            answer += "IX";
        }

        num /= 10;
        int tens = num % 10;
        string tensString = "";

        if (tens <= 3) {
            tensString.append(tens, 'X');
        } else if (tens == 4) {
            tensString += "XL";
        } else if (tens < 9) {
            tensString += "L";
            tensString.append(tens - 5, 'X');
        } else if (tens == 9) {
            tensString += "XC";
        }

        answer = tensString + answer;

        num /= 10;
        int hundreds = num % 10;
        string hundredsString = "";

        if (hundreds <= 3) {
            hundredsString.append(hundreds, 'C');
        } else if (hundreds == 4) {
            hundredsString += "CD";
        } else if (hundreds < 9) {
            hundredsString += "D";
            hundredsString.append(hundreds - 5, 'C');
        } else if (hundreds == 9) {
            hundredsString += "CM";
        }

        answer = hundredsString + answer;

        num /= 10;

        string thousands = "";
        thousands.append(num % 10, 'M');

        return thousands + answer;
    }
};

int main() {
    Solution s;

    for (int i = 0; i <= 10; ++i) {
        cout << s.intToRoman(i) << endl;
    }

    for (int i = 8; i < 3999; i *= 2) {
        cout << s.intToRoman(i) << endl;
    }
}