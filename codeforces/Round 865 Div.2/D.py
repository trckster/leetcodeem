import sys


def get_length(i, j):
    print(f'? {i + 1} {j + 1}')
    sys.stdout.flush()
    return int(input())


t = int(input())

for _ in range(t):
    n = int(input())

    chain = []

    i = 1
    for __ in range(n // 2):
        chain.append(n - i + 1)
        chain.append(i)
        i += 1
    if n % 2 == 1:
        chain.append(i)

    print(f'+ {n}')
    sys.stdout.flush()
    input()
    print(f'+ {n + 1}')
    sys.stdout.flush()
    input()

    destinations = [-1]

    for j in range(1, n):
        destinations.append(get_length(0, j))

    answer1 = [0 for i in range(n)]
    answer2 = [0 for i in range(n)]

    max_destination = max(destinations)
    peak_index = destinations.index(max_destination)

    answer1[0] = chain[max_destination]
    answer1[peak_index] = chain[0]
    answer2[0] = chain[-max_destination - 1]
    answer2[peak_index] = chain[-1]

    i = 1
    while i < n:
        if i == peak_index:
            i += 1
            continue

        length_between_peak_and_i = get_length(peak_index, i)
        answer1[i] = chain[length_between_peak_and_i]
        answer2[i] = chain[-length_between_peak_and_i - 1]

        i += 1

    answer1_s = ' '.join(map(str, answer1))
    answer2_s = ' '.join(map(str, answer2))
    print(f'! {answer1_s} {answer2_s}')
    sys.stdout.flush()
    ____ = input()
