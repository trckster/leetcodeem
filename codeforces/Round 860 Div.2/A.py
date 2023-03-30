t = int(input())


def is_possible(a, b) -> bool:
    biggest = max(max(a), max(b))

    if a[-1] != biggest and b[-1] != biggest:
        return False

    for i in range(len(a)):
        if a[i] < b[i]:
            a[i], b[i] = b[i], a[i]

    return max(b) == b[-1]


for _ in range(t):
    n = int(input())

    a = list(map(int, input().split()))
    b = list(map(int, input().split()))

    print('Yes' if is_possible(a, b) else 'No')
