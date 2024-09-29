Picking Numbers
================================

Given an array of integers, find the longest subarray where the absolute difference between any two elements is less than or equal to `1`.

### Example 1
```
[in]
4 6 5 3 3 1

[out]
3
```

We choose the following multiset of integers from the array: `{4, 3, 3}`. Each pair in the multiset has an absolute difference lowest or equal `1`. (ex: 4 - 3 = 1 and 3 - 3 = 0), so we print the number of chosen integer 3, as our answer. 
