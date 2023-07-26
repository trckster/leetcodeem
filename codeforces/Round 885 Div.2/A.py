def get_ints():
    return map(int, input().split())


def process():
    n, m, k = get_ints()
    v_x, v_y = get_ints()
    v_color = (v_x + v_y) % 2

    runs = True

    for _ in range(k):
        x, y = get_ints()

        if v_color == (x + y) % 2:
            runs = False

    print('Yes' if runs else 'No')


t = int(input())

for _ in range(t):
    process()
