## Leetcode 1337. The K Weakest Rows in a Matrix
---
You are given an `m x n` binary matrix mat of `1`'s (representing soldiers) and `0`'s (representing civilians). The soldiers are positioned in front of the civilians. That is, all the `1`'s will appear to the left of all the 0's in each row.

A row `i` is weaker than a row `j` if one of the following is true:

- The number of soldiers in row `i` is less than the number of soldiers in row `j`.
- Both rows have the same number of soldiers and `i < j`.


Return the indices of the `k` weakest rows in the matrix ordered from weakest to strongest.

---
### How to slove this problem?
1. `Calculate` the number of soldiers in each rows in order to sort the power of each rows.
2. `Sort` the matrix, in which the power of each row is.
3. `Like` the matrix with row number with the power of matrix while sorting.
4. `insert` the `k` weakest rows in the vector.
5. `return`
---
```
class Solution {
public:
    vector<int> kWeakestRows(vector<vector<int>>& mat, int k) {
        int pos = mat[0].size(); // get the end position of each row.
        for(int i=0;i<mat.size();i++)
            mat[i].push_back(i); // insert the each row's number at the end of each row.
        sort(mat.begin(),mat.end()); // sorting the 2D vector using sort fuction.  
        vector<int> ans;
        for(int i=0;i<k;i++)
            ans.push_back(mat[i][pos]); // put the sorted priority of each row to the k weakest row.
        return ans;
    }
};
```
---
## Reference
1. [Sort](https://www.geeksforgeeks.org/sorting-a-vector-in-c/)
2. [Vector](https://shengyu7697.github.io/std-vector/)
3. [Vedio](https://www.youtube.com/watch?v=GICnxgx1lXw)
