class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        int size = nums.size(); // to know how many time we need to run to find the ans.
        vector<int> sum_num;
        for(int i=0;i<size-1;i++)
            for(int j=i+1;j<size;j++)
                if(nums[i]+nums[j]==target) {
                    sum_num.push_back(i);
                    sum_num.push_back(j);
                    break;
                }
        //when it meet the target push two fit element in vector and reutrn this vector
        return sum_num;
    }
};
