guessing_height = []


def ceil(a, b):
    return -(-a // b)


def count_days(h, a, b):
    last_night = max(0, h - a)
    return max(1, ceil(last_night, a - b) + 1)


def add_statement(a, b, n):
    if n == 1:
        h1 = 0
        h2 = a
    else:
        h1 = max(0, (n - 2) * (a - b) + a)
        h2 = max(a, h1 + a - b)

    h1 += 1

    if h1 > guessing_height[1] or h2 < guessing_height[0]:
        return 0

    guessing_height[0] = max(h1, guessing_height[0])
    guessing_height[1] = min(h2, guessing_height[1])

    return 1


def how_long(a, b):
    min_count = count_days(guessing_height[0], a, b)
    max_count = count_days(guessing_height[1], a, b)

    if min_count == max_count:
        return min_count
    else:
        return -1


t = int(input())

for _ in range(t):
    q = int(input())

    guessing_height = [1, 2 ** 63 - 1]

    answer = []
    for __ in range(q):
        event = input().split()

        if event[0] == '1':
            answer.append(str(add_statement(int(event[1]), int(event[2]), int(event[3]))))
        else:
            answer.append(str(how_long(int(event[1]), int(event[2]))))

    print(' '.join(answer))
