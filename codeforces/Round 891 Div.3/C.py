t = int(input())

for _ in range(t):
    n = int(input())
    arr = list(map(int, input().split()))

    arr.sort()

    answer = []
    jump = n - 1
    i = 0
    while jump > 0:
        answer.append(arr[i])

        i += jump
        jump -= 1

    answer.append(answer[-1])

    print(' '.join(map(str, answer)))
