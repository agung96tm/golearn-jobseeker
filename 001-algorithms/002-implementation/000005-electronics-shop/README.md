Electronics Shop
===============================

A person wants to determine the most expensive computer keyboard and USB drive that can be purchased with a give budget. Given price lists for keyboards and USB drives and a budget, find the cost to buy them. If it is not possible to buy both items, return `-1`.


### Example 1:
```
b = 60
keyboards = [40, 50, 60]
drives = [5, 8, 12]
```

The person can buy a `40 Keyboard + 12 Drive = 52`, or a `50 Keyboard + 8 Drive = 58`. Choose the latter as the more expensive option and return `58`.


### Example 2:
```
b = 5
keyboards = [4]
drives = [5]
```

There is no way to buy one keyboard and one USB drive because `4 + 5 = 9`, so return `-1`.