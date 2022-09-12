
from sys import flags


def myPow(x: float, n: int) -> float:
    if n >= 0:
        return quickMul(x, n)
    return 1.0 / quickMul(x,-n)

def quickMul(x: float, n: int) -> float:
    if n == 0:
        return 1
    y = quickMul(x, n//2)
    if n %2 == 0:
        return y*y
    return y*y*x