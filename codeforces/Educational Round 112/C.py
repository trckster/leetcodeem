t = int(input())

for _ in range(t):
    m = int(input())
    first = list(map(int, input().split()))
    second = list(map(int, input().split()))

    first_sum = sum(first)

    left_abandoned = 0
    right_abandoned = first_sum - first[0]
    alice_best = right_abandoned
    i = 0
    while True:
        bob_best = max(left_abandoned, right_abandoned)
        alice_best = min(alice_best, bob_best)

        if i >= len(first) - 1:
            break

        left_abandoned += second[i]
        right_abandoned -= first[i + 1]

        i += 1

    print(alice_best)
