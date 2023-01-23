class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        size_t len = 0; // the new location was needed to return
        for(int i=0;i<nums.size();i++)
            if(nums[i]!=val)
                nums[len++]=nums[i]; // insert the element in original sequencial array
        return len;
    }
};