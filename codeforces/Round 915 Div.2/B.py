from math import ceil

def yes():
    print('YES')


def no():
    print('NO')


def solve():
    n = int(input())

    nodes = [[] for _ in range(n)]
    for _ in range(n - 1):
        u, v = map(int, input().split())
        nodes[u - 1].append(v - 1)
        nodes[v - 1].append(u - 1)

    leafs = 0
    for connected in nodes:
        if len(connected) == 1:
            leafs += 1
    print(ceil(leafs / 2))


t = int(input())
for _ in range(t):
    solve()
