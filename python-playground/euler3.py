# The prime factors of 13195 are 5, 7, 13 and 29.

# What is the largest prime factor of the number 600851475143 ?

from math import trunc, sqrt

PRIMES = []
NOT_PRIMES = []


def is_prime(n):
    if n in NOT_PRIMES: return False
    if n in PRIMES: return True
    for i in range(2, trunc(sqrt(n)) + 1):
        if n % i == 0:
            NOT_PRIMES.append(n)
            return False
    PRIMES.append(n)
    return True


def primes(n):
    p = []
    for i in range(2, n):
        if is_prime(i):
            p.append(i)
    return p


def prime_factors(n):
    f = []
    i = 2
    while i <= n:
        if is_prime(i) and n % i == 0:
            f.append(i)
            n = n / i
            i = 2
        else:
            i += 1
    return f


print(prime_factors(600851475143))
