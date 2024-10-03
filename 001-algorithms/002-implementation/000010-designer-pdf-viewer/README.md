Designer PDF Viewer
============================

There is a list of `26` character heights aligned by index to their letters. For example, 'a' is at index `0` and 'z' is at index `25`. There will also be a string. Using the letter heights given, determine the area of the rectangle highlight in `mm2` assuming all letters are `1mm` wide.


### Example 1
```
[in]
    // heights
    1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5
    
    // word
    torn  

[out]
    8
```
The heights are `t = 2`, `0 = 1`, `r = 1` and `n = 1`. The tallest letter is `2` high and there are `4` letters. The hightlighted area will be `2 * 4 = 8mm2` so the answer is `8`.



### Example 2
```
[in]
    // heights
    1 3 1 3 1 4 1 3 2 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5 5
    
    // word
    abc  

[out]
    9
```
The heights are `a = 1`, `b = 3`, and `c = 1`. The tallest letter is `3` high and there are `3` letters. The hightlighted area will be `3 * 3 = 9mm2` so the answer is `9`.