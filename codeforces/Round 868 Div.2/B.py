def solve():
    n, k = map(int, input().split())
    arr = list(map(int, input().split()))

    invalids = []

    for i in range(len(arr)):
        real_place = i
        valid_place = arr[i] - 1
        distance = abs(real_place - valid_place)

        if distance % k == 0:
            continue

        invalids.append(real_place)

    if len(invalids) == 0:
        print(0)
    elif len(invalids) == 2:
        first_invalid = invalids[0]
        second_invalid = invalids[1]

        arr[first_invalid], arr[second_invalid] = arr[second_invalid], arr[first_invalid]

        distance_1 = abs(arr[first_invalid] - 1 - first_invalid)
        distance_2 = abs(arr[second_invalid] - 1 - second_invalid)

        if distance_1 % k == 0 and distance_2 % k == 0:
            print(1)
        else:
            print(-1)
    else:
        print(-1)


t = int(input())
for _ in range(t):
    solve()
