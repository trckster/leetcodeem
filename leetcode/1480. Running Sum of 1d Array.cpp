class Solution {
public:
    vector<int> runningSum(vector<int>& nums) {
        vector<int> answer;
        answer.push_back(nums[0]);
        
        for (int j = 1; j < nums.size(); j++) {
            answer.push_back(answer[j-1] + nums[j]);
        }
        
        return answer;
    }
};
