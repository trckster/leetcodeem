def get_best_mex(arr):
    return len(arr) + arr.count(0)


def get_weight(arr):
    final_sum = 0

    l = 0
    while l < len(arr):
        r = l + 1
        while r < len(arr) + 1:
            current_sub_section = arr[l:r]
            final_sum += get_best_mex(current_sub_section)

            r += 1

        l += 1

    return final_sum


t = int(input())

for i in range(t):
    n = int(input())
    arr = list(map(int, input().split()))

    print(get_weight(arr))
