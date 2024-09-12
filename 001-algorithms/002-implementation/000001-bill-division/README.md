## Bill Division
-----------------------

```
[in]  
1
3 10 2 9
12 

[out] 
5
```
Skip position 1 `10` then calculate others as total (3 + 2 + 9 = 14). Split total with `2` (two persons) (14 / 2 = 7). Then `12` as bill, minus with `7` (12 - 7 = 5) result become `5`.


```
[in]  
1
3 10 2 9
7

[out] 
Bon Appetit
```
When result is `0` then print `Bon Appetit`.