Climbing the Leaderboard
================================

An arcade game player wants to climb to the top of the leaderboard and track their ranking. The game uses Dense Ranking, so its leaderboard works like this:

The player with the highest score is ranked number `1` on the leaderboard.
Players who have equal scores receive the same ranking number, and the next player(s) receive the immediately following ranking number.

### Example 1
```
[in]
100, 90, 90, 80 // ranked
70, 80, 105 // players

[out]
4, 3, 1
```

The ranked players will have ranks `1`, `2`, `2`, `1`, and , respectively. If the player's scores are `70`, `80` and `105`, their rankings after each game are `4`, `3` and `1`. Return `[4, 3, 1]`.