Counting Valleys
===============================

An avid hiker keeps meticulous records of their hikes. During the last hike that took exactly `steps` steps, for every step it was noted if it was an uphill,`U` , or a downhill, `D` step. Hikes always start and end at sea level, and each step up or down represents a `1` unit change in altitude. We define the following terms:

* A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a step down to sea level.
* A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level.

Given the sequence of up and down steps during a hike, find and print the number of valleys walked through.

### Example:
```
[in]
8
DDUUUUDD

[out]
0
```

### Explain:
If we represent `_` as sea level, a step up as `/`, and a step down as ` \ `, the hike can be drawn as:
```

        /\
_    _ /  \ _
 \  /
  \/
```

The hiker first enters a valley `2` units deep. Then they climb out `valleyCount = 1` and up onto a mountain `2` units high. Finally, the hiker returns to sea level `mountainCount = 1` and ends the hike.