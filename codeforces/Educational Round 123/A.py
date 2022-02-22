t = int(input())

doors = 'RGB'

m = {
    'R': 'r',
    'G': 'g',
    'B': 'b'
}

for task in range(t):
    keys = ''
    answer = 'YES'
    s = input()
    for c in s:
        if c in doors:
            if m[c] not in keys:
                answer = 'NO'
        else:
            keys += c

    print(answer)
