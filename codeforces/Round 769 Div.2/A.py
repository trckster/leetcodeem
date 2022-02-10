n = int(input())

for i in range(n):
    s_len = int(input())
    s = input()

    if s_len >= 3:
        print('NO')
        continue
    elif s_len == 1:
        print('YES')
    elif s == '10' or s == '01':
        print('YES')
    else:
        print('NO')
