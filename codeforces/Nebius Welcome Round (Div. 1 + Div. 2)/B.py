t = int(input())

for _ in range(t):
    n, k, d, w = map(int, input().split())
    t = list(map(int, input().split()))

    package = 0

    i = 0
    while i < len(t):
        package += 1
        inject_time = t[i] + w
        i += 1
        last_time_to_inject = inject_time + d
        doses_count = k - 1

        while doses_count > 0 and i < len(t):

            if t[i] > last_time_to_inject:
                break

            doses_count -= 1
            i += 1

    print(package)
