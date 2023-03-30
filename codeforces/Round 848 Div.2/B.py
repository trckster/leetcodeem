t = int(input())

for _ in range(t):
    n, m, d = map(int, input().split())
    p = list(map(int, input().split()))
    a = list(map(int, input().split()))

    value_to_position = {}
    for i in range(len(p)):
        value_to_position[p[i]] = i

    answer = 1000000

    for i in range(len(a) - 1):
        a1 = a[i]
        a2 = a[i + 1]

        pos1 = value_to_position[a1]
        pos2 = value_to_position[a2]

        distance = pos2 - pos1

        distance_to_separation = d - distance + 1

        if distance + distance_to_separation >= n:
            distance_to_separation = 1000000

        best_for_pair = min(distance, distance_to_separation)
        answer = min(answer, best_for_pair)

    print(max(answer, 0))
