Beautiful Days at the Movies
============================

Lily likes to play games with integers. She has created a new game where she determines the difference between a number and its reverse. 
For instance, given the number `12`, its reverse is `21`. Their difference is `9`. The number `120` reversed is `21`, and their difference is `99`.

She decides to apply her game to decision making. She will look at a numbered range of days and will only go to a movie on a beautiful day.

### Explain 1
```
[in]
20 23 6

[out]
2
```

Lily may go to the movies on days `20`, `21`, `22`, and `23`. We perform the following calculations to determine which days are beautiful:

* Day `20` is beautiful because the following evaluates to a whole number:
  * `| 20 - 02 | / 6` = `18 / 6` = `3`
* Day `21` is not beautiful because the following doesn't evaluate to a whole number:
  * `| 21 - 12 | / 6` = `9 / 6` = `1.5`
* Day `22` is beautiful because the following evaluates to a whole number:
  * `| 22 - 22 | / 6` = `0 / 6` = `0`
* Day `23` is not beautiful because the following doesn't evaluate to a whole number:
  * `| 23 - 32 | / 6` = `9 / 6` = `1.5`

Only two days, `20` and `22`, in this interval are beautiful. Thus, we print `2` as our answer.