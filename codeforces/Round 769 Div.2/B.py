def get_closest_power_of_2(n):
    x = 1
    while 2 * x < n:
        x *= 2
    return x


def print_sequence(n):
    closest_power = get_closest_power_of_2(n)

    for i in range(1, closest_power):
        print(i, end=' ')

    print(0, end=' ')

    for i in range(closest_power, n):
        print(i, end=' ')

    print()


t = int(input())

for i in range(t):
    n = int(input())
    print_sequence(n)
