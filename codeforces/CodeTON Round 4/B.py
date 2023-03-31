def process(k):
    if k % 2 == 0:
        print(-1)
        return

    binary = bin(k)[2:-1]

    print(len(binary))

    for i in binary:
        print(2 if int(i) % 2 else 1, end=' ')
    print()


t = int(input())

for _ in range(t):
    n = int(input())

    process(n)
