from itertools import count

D = {1: 1, 2: 2}


def fib(n):
    if n < 1:
        raise RuntimeError("n should not be less than one: %d" % n)
    if n in D:
        return D[n]
    res = fib(n - 2) + fib(n - 1)
    D[n] = res
    return res


assert fib(1) == 1
assert fib(2) == 2
assert fib(3) == 3
assert fib(4) == 5

s = 0
for f in count(start=1):
    v = fib(f)
    if v > 4000000:
        break
    elif v % 2 == 0:
        s += v

print(s)

# assert res == [1, 2, 3, 5, 8, 13, 21, 34, 55, 89]

print('done')
