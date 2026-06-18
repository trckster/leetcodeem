import sys
input = sys.stdin.readline
MOD = 10**9 + 7
INF = 10**30


def ints():
    return list(map(int, input().split()))


def ii():
    return int(input())


def si():
    return input().strip()


def yes(ok=True):
    print("YES" if ok else "NO")


def no():
    print("NO")


def ceil_div(a, b):
    return -(-a // b)


def solve_case():
    _ = ii()
    a = ints()

    m = sys.maxsize
    answer = 0
    for t in a:
        m = min(m, t)
        answer += m
    print(answer)


def main():
    t = ii()
    for _ in range(t):
        solve_case()


if __name__ == "__main__":
    main()
