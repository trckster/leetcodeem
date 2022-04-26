t = int(input())

for _t in range(t):
    s = input()

    A_count = B_count = 0

    for c in s:
        if c == 'A':
            A_count += 1
        else:
            A_count -= 1

        if A_count < 0:
            break

    if A_count < 0 or s.endswith('A'):
        print('NO')
    else:
        print('YES')
