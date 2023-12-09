def yes():
    print('YES')


def no():
    print('NO')


def solve():
    a, b = map(int, input().split())
    xk, yk = map(int, input().split())
    xq, yq = map(int, input().split())

    k = set()
    k.add('{}-{}'.format(xk + a, yk + b))
    k.add('{}-{}'.format(xk + a, yk - b))
    k.add('{}-{}'.format(xk - a, yk + b))
    k.add('{}-{}'.format(xk - a, yk - b))
    k.add('{}-{}'.format(xk + b, yk + a))
    k.add('{}-{}'.format(xk + b, yk - a))
    k.add('{}-{}'.format(xk - b, yk + a))
    k.add('{}-{}'.format(xk - b, yk - a))

    q = set()
    q.add('{}-{}'.format(xq + a, yq + b))
    q.add('{}-{}'.format(xq + a, yq - b))
    q.add('{}-{}'.format(xq - a, yq + b))
    q.add('{}-{}'.format(xq - a, yq - b))
    q.add('{}-{}'.format(xq + b, yq + a))
    q.add('{}-{}'.format(xq + b, yq - a))
    q.add('{}-{}'.format(xq - b, yq + a))
    q.add('{}-{}'.format(xq - b, yq - a))

    print(len(k.intersection(q)))


t = int(input())
for _ in range(t):
    solve()
