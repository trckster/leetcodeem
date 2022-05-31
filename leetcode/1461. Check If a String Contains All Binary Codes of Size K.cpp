#include <iostream>
#include <set>

using namespace std;

class Solution {
public:
    bool hasAllCodes(string s, int k) {
        int maxValue = (1 << k) - 1;

        int sLength = s.length();

        if (maxValue > sLength - k + 1) {
            return false;
        }

        int item = 0, i = 0;
        for (; i < k; ++i) {
            item <<= 1;

            if (s[i] == '1') {
                item++;
            }
        }

        set<int> set;
        set.insert(item);

        for (; i < s.length(); ++i) {
            item <<= 1;
            if (s[i] == '1') {
                item++;
            }

            item &= maxValue;

            set.insert(item);
        }

        for (i = 0; i <= maxValue; ++i) {
            if (!set.count(i)) {
                return false;
            }
        }

        return true;
    }
};

int main() {
    Solution s;

    cout << (s.hasAllCodes("1100101010", 1) ? 'T' : 'F') << endl;
    cout << (s.hasAllCodes("00110", 2) ? 'T' : 'F') << endl;
    cout << (s.hasAllCodes("1100101010", 2) ? 'T' : 'F') << endl;
    cout << (s.hasAllCodes("00010110010110111", 3) ? 'T' : 'F') << endl;
    cout << (s.hasAllCodes("00001011001011011100010110011011011100010110010101110001001100101101111", 4) ? 'T' : 'F')
         << endl;
}