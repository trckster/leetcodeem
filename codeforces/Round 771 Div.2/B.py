t = int(input())

for i in range(t):
    n = int(input())
    arr = list(map(int, input().split()))

    lastEven = 0
    lastOdd = 0
    answer = 'Yes'

    for k in arr:
        if k % 2:
            if lastOdd > k:
                answer = 'No'
                break
            else:
                lastOdd = k
        else:
            if lastEven > k:
                answer = 'No'
                break
            else:
                lastEven = k

    print(answer)