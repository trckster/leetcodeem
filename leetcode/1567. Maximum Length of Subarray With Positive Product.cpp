#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
    bool hasPositiveProduct(vector<int> arr) {
        bool hasPositiveProduct = true;

        for (auto item: arr) {
            if (item < 0) {
                hasPositiveProduct = !hasPositiveProduct;
            }
        }

        return hasPositiveProduct;
    }

    int getMaxLenOfArrayWithoutNulls(vector<int> arr) {
        if (hasPositiveProduct(arr)) {
            return arr.size();
        }

        for (int i = 0; i < arr.size(); i++) {
            if (arr[i] < 0 || arr[arr.size() - 1 - i] < 0) {
                return arr.size() - i - 1;
            }
        }

        return 0;
    }

public:
    int getMaxLen(vector<int> &nums) {
        vector <vector<int>> subArrays;

        nums.insert(nums.begin(), 0);

        for (int i = 1; i < nums.size(); ++i) {
            if (!nums[i]) {
                continue;
            }

            if (!nums[i - 1]) {
                vector<int> subArray;
                subArrays.push_back(subArray);
            }

            subArrays[subArrays.size() - 1].push_back(nums[i]);
        }

        int answer = 0;

        for (auto arr: subArrays) {
            answer = max(answer, getMaxLenOfArrayWithoutNulls(arr));
        }

        return answer;
    }
};

int main() {
    Solution s;

    vector<int> arr{1, -2, -3, 4};
    cout << s.getMaxLen(arr) << endl;
    arr = {0, 1, -2, -3, -4};
    cout << s.getMaxLen(arr) << endl;
    arr = {-1, -2, -3, 0, 1};
    cout << s.getMaxLen(arr) << endl;
    arr = {1, 2, 3, 5, -6, 4, 0, 10};
    cout << s.getMaxLen(arr) << endl;
}