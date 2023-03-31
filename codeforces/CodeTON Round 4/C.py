def process():
    pass


t = int(input())

for _ in range(t):
    n, c, d = map(int, input().split())
    arr = list(map(int, input().split()))

    unique = list(set(arr))

    minimal_cost = c * (len(arr) - len(unique))

    unique.sort()

    if unique[0] != 1:
        unique.insert(0, 1)
        minimal_cost += d

    start = None

    for i in range(len(unique)):
        if i + 1 != unique[i]:
            start = i - 1
            break

    if start is None:
        print(minimal_cost)
        continue

    final_cost = 2 ** 63 - 1

    for i in range(start, len(unique)):
        to_remove = len(unique) - i - 1

        to_add = unique[i] - (i + 1)

        cost = c * to_remove + d * to_add

        final_cost = min(cost, final_cost)

    print(final_cost + minimal_cost)
