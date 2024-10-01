The Hurdle Race

A video player plays a game in which the character competes in a hurdle race. Hurdles are of varying heights, and the characters have a maximum height they can jump. There is a magic potion they can take that will increase their maximum jump height by `1` unit for each dose. How many doses of the potion must the character take to be able to jump all of the hurdles. If the character can already clear all of the hurdles, return `0`.

### Example 1
```
[in]
1 2 3 3 2   // heights
1           // character_jump

[out]
2
```

The character can jump `1` unit high initially and must take `3 - 1 = "2"` doses of potion to be able to jump all of the hurdles.


### Example 2
```
[in]
1 6 3 5 2   // heights
4           // character_jump

[out]
2
```

The character can jump `4` unit high initially and must take `6 - 4 = "2"`  doses of potion to be able to jump all of the hurdles.