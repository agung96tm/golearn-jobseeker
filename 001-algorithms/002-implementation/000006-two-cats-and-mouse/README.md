Two cats and a mouse
================================

Two cats and a mouse are at various positions on a line. You will be given their starting positions. Your task is to determine which cat will reach the mouse first, assuming the mouse does not move and the cats travel at equal speed. If the cats arrive at the same time, the mouse will be allowed to move and it will escape while they fight.


* If cat `A` catches the mouse first, print Cat A.
* If cat `B` catches the mouse first, print Cat B.
* If both cats reach the mouse at the same time, print Mouse `C` as the two cats fight and mouse escapes.

### Example 1
```
[in]
a = 2
b = 5
c = 4

[out]
Cat B
```

The cats are at positions `2` (Cat A) and `5` (Cat B), and the mouse is at position `4`. Cat B, at position `5` will arrive first since it is only `1` unit away while the other is `2` units away. Return 'Cat B'.


### Example 2
```
[in]
a = 1 
b = 3
c = 2

```

The cats are at positions `1` (Cat A) and `3` (Cat B), and the mouse is at position `2`. cats  and  reach mouse `2` at the exact same time, we print 'Mouse C' on a new line.