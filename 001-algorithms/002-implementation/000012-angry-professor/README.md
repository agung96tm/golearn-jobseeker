Angry Professor
============================

A Discrete Mathematics professor has a class of students. Frustrated with their lack of discipline, 
the professor decides to cancel class if fewer than some number of students are present when class starts. 
Arrival times go from on time (`arrivalTime <= 0`) to arrived late (`arrivalTime > 0`).

### Explain 1
```
a = -2, -1, 0, 1, 2
k = 3
```

The first `3` students arrived on. The last `2` were late. The threshold is `3` students, so class will go on. Return YES.


### Explain 2
```
a = 0, -1, 2, 1
k = 2
```

The first `2` students arrived on. The last `2` were late. The threshold is `2` students, so class will go on. Return NO.

