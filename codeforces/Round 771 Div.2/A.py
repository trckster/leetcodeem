def print_permutation(p):
    for i in p:
        print(i, end=' ')
    print()


t = int(input())

for i in range(t):
    n = int(input())
    permutation = list(map(int, input().split()))

    l = None
    j = 0
    while j < n:
        if permutation[j] != j + 1:
            l = j
            break
        j += 1

    if l is not None:
        r = permutation.index(j + 1)

        new_permutation = permutation[:l] + permutation[l:r + 1][::-1] + permutation[r + 1:]
        print_permutation(new_permutation)
    else:
        print_permutation(permutation)
