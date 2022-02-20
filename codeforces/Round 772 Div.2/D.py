fibs = {
    1: 1,
    2: 1,
}

fib_sums = {
    1: 1,
}

LIMIT = 10 ** 9 + 7

max_p = 2 * (10 ** 5) + 1

i = 3
while i < max_p:
    fibs[i] = (fibs[i - 1] + fibs[i - 2]) % LIMIT
    i += 1

i = 2
while i < max_p:
    fib_sums[i] = (fib_sums[i - 1] + fibs[i]) % LIMIT
    i += 1

numbers = set()


def unique(number):
    while number:
        if number & 1 == 1:
            number >>= 1
        elif number & 3 == 0:
            number >>= 2
        else:
            return True

        if number in numbers:
            return False

    return True


def get_msb(n):
    result = 0
    while n:
        result += 1
        n >>= 1
    return result


n, p = map(int, input().split())
arr = list(map(int, input().split()))
arr = sorted(arr)

answer = 0

for i in arr:
    if unique(i):
        numbers.add(i)

        msb = get_msb(i)
        jumps = p - msb + 1

        if jumps <= 0:
            continue

        answer = (answer + fib_sums[jumps]) % LIMIT

print(answer)
