def process():
    n = int(input())
    arr = list(map(int, input().split()))

    remember = []
    odd = []
    even = []
    for k in arr:
        if k % 2:
            odd.append(k)
            remember.append(True)
        else:
            even.append(k)
            remember.append(False)

    odd.sort()
    even.sort()

    odd_i = 0
    even_i = 0
    for i in range(len(arr)):
        if remember[i]:
            remember[i] = odd[odd_i]
            odd_i += 1
        else:
            remember[i] = even[even_i]
            even_i += 1

    is_sorted = all(remember[i] <= remember[i + 1] for i in range(len(remember) - 1))

    print('Yes' if is_sorted else 'No')


t = int(input())
for _ in range(t):
    process()
