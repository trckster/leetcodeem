t = int(input())

for _ in range(t):
    a, b = map(int, input().split())

    a_abs = abs(a)
    b_abs = abs(b)
    wait = max(0, abs(a_abs - b_abs) - 1)

    answer = min(a_abs, b_abs) * 2 + max(a_abs, b_abs) - min(a_abs, b_abs) + wait

    print(answer)
