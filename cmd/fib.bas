2030 DIM f(51)
2040 LET f(0) = 0
2050 LET f(1) = 1
2060 FOR n = 2 TO 45
2080 LET f(n) = f(n-1) + f(n-2)
2090 NEXT n
