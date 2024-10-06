Utopian Tree
============================

The Utopian Tree goes through 2 cycles of growth every year. Each spring, it doubles in height. Each summer, its height increases by 1 meter.

A Utopian Tree sapling with a height of 1 meter is planted at the onset of spring. How tall will the tree be after  growth cycles?

For example, if the number of growth cycles is `n = 5`, the calculations are as follows:

### Explain
```
Period  Height
0          1
1          2
2          3
3          6
4          7
5          14
```

---

### Example
```
[in]
0
1
4

[out]
1
2
7
```
There are 3 test cases.

In the first case (`n = 0`), the initial height (`1`) of the tree remains unchanged.

In the second case (`n = 1`), the tree doubles in height and is `2` meters tall after the spring cycle.

In the third case (`n = 4`), the tree doubles its height in spring (`n = 1`, `height = 2`), then grows a meter in summer (`n = 2`, `height = 3`), then doubles after the next spring (`n = 3`, `height = 6`), and grows another meter after summer (`n = 4`, `height = 7`). Thus, at the end of 4 cycles, its height is `7` meters.