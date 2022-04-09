MOD = 998244353

def solve(n):
    if n % 2:
        return 0

    answer = 1

    for x in range(1, n // 2 + 1):
        answer *= x
        answer %= MOD

    answer = (answer ** 2) % MOD

    return answer


t = int(input())

for i in range(t):
    n = int(input())
    print(solve(n))
