# digits-helper

A solve for the NYTimes Digits game!
https://www.nytimes.com/games/digits

To run from source

1. Clone the repo
2. Run with go run

This does not try to find the 'best' solution or list all solutions.
It just lists the first one that we find.

## Examples

```
❯ go run . 59 2 3 5 11 15 25     
*** SOLUTION ***
25 + 15
40 + 11
51 + 5
56 + 3
```

```
❯ go run . 133 4 5 8 11 15 20
*** SOLUTION ***
20 + 15
35 - 11
24 + 8
32 * 4
128 + 5
```