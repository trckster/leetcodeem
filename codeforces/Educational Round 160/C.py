def yes():
    print('YES')


def no():
    print('NO')


powers = [2 ** i for i in range(30)][::-1]


def split(x):
    parts = []
    for power in powers:
        if x >= power:
            parts.append(power)
            x -= power
    return parts[::-1]


p_2_30 = 2 ** 30


class MS:
    def __init__(self):
        self.storage = {}
        for power in powers:
            self.storage[power] = 0

    def add(self, x):
        x = 2 ** x
        if x in self.storage:
            self.storage[x] += 1
        else:
            self.storage[x] = 1

    def get(self, w):
        parts = split(w)
        taken = {}

        luggage = 0
        current_power = 1
        for p in parts:
            while current_power <= p_2_30:
                if current_power == p_2_30:
                    return False

                power_cap = self.storage[current_power] * current_power
                if power_cap + luggage >= p:
                    luggage += power_cap - p
                    current_power *= 2
                    break
                else:
                    luggage += power_cap
                    current_power *= 2

                if current_power > p:
                    return False
        return True


t = int(input())
ms = MS()
for _ in range(t):
    t, v = map(int, input().split())

    if t == 1:
        ms.add(v)
    else:
        yes() if ms.get(v) else no()
