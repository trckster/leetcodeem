from math import ceil


def solve(team1, team2):
    if not team2:
        return 'R' * team1

    surpass_magnitude = ceil(team1 / (team2 + 1))

    answer = 'R' * surpass_magnitude + 'B'

    return answer + solve(team1 - surpass_magnitude, team2 - 1)


t = int(input())

for s in range(t):
    n, r, b = map(int, input().split())

    print(solve(r, b))
