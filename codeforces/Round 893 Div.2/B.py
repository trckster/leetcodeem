def solve():
    n, m, d = map(int, input().split())
    arr = list(map(int, input().split()))

    arr.sort()

    if arr[0] != 1:
        arr.insert(0, 1)
    arr.append(n + 1)

    important_shops_count = 0
    stops_with_all_shops = 0

    i = 1
    while i < len(arr) - 1:
        before_this_shop = 1 + (arr[i] - arr[i - 1] - 1) // d
        stops_with_i_shop = before_this_shop + 1 + (arr[i + 1] - arr[i] - 1) // d
        stops_without_i_shop = 1 + (arr[i + 1] - arr[i - 1] - 1) // d

        stops_with_all_shops += before_this_shop

        if stops_with_i_shop != stops_without_i_shop:
            important_shops_count += 1

        i += 1

    stops_with_all_shops += 1 + (arr[-1] - arr[-2] - 1) // d

    if important_shops_count > 0:
        stops_with_all_shops -= 1

    if important_shops_count == 0:
        important_shops_count = m

    print(stops_with_all_shops, important_shops_count)


t = int(input())

for _ in range(t):
    solve()
