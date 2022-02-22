def find_max_subarray(arr):
    current_max = arr[0]
    global_max = arr[0]
    end = 0

    for i in range(1, len(arr)):
        current_max = max(arr[i], arr[i] + current_max)

        if current_max > global_max:
            global_max = current_max
            end = i

    start = end

    while start >= 0:
        global_max -= arr[start]

        if global_max == 0:
            break

        start -= 1

    return start, end, sum(arr[start:end + 1])


def scan(arr, number, numbers_available):
    best_sum = initial_sum = sum(arr[0:numbers_available])

    i = 0
    while numbers_available + i < len(arr):
        initial_sum -= arr[i]
        initial_sum += arr[i + numbers_available]

        if best_sum < initial_sum:
            best_sum = initial_sum

        i += 1

    return best_sum + number * numbers_available


def print_array(arr):
    for i in arr:
        print(i, end=' ')
    print()


def all_less_or_equal(arr, value):
    return all(i <= value for i in arr)


def get_max_index(arr):
    m = arr[0]
    m_i = 0

    for i in range(len(arr)):
        if m < arr[i]:
            m = arr[i]
            m_i = i

    return m_i


t = int(input())

for task in range(t):
    n, x = map(int, input().split())
    data = list(map(int, input().split()))

    k = 0
    l = r = None
    answer = []
    if all_less_or_equal(data, 0):
        answer.append(0)
        k += 1

        max_index = get_max_index(data)

        if data[max_index] + x <= 0:
            answer = [0 for i in range(n)]
            k += n
        else:
            answer.append(data[max_index] + x)
            k += 1
            l = r = max_index
    else:
        l, r, best_result = find_max_subarray(data)

        answer.append(best_result)
        k += 1

        for j in range(r - l + 1):
            answer.append(answer[k - 1] + x)
            k += 1

    while k <= n:
        best_local_result = scan(data, x, k)
        answer.append(max(best_local_result, answer[k - 1]))

        k += 1

    if t == 1:
        print(answer[508], answer[509], answer[510], answer[511])

    print_array(answer)
