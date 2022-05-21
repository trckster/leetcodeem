class Solution:
    def __init__(self):
        self.coins = []
        self.cache = {}

    def best_way_to_reach(self, target):
        if target in self.cache:
            return self.cache[target]

        answer = 10001
        for coin in self.coins:
            if target - coin < 0:
                continue

            if target - coin == 0:
                return 1

            result = self.best_way_to_reach(target - coin)

            if result != -1 and result + 1 < answer:
                answer = result + 1

        self.cache[target] = -1 if answer == 10001 else answer
        return self.cache[target]

    def set_coins(self, coins, target):
        suitable_coins = []
        for coin in coins:
            if coin <= target:
                suitable_coins.append(coin)

        self.coins = suitable_coins

    def coinChange(self, coins, amount):
        if amount == 0:
            return 0

        self.cache = {}
        self.set_coins(coins, amount)

        return self.best_way_to_reach(amount)


s = Solution()
print(s.coinChange([186, 419, 83, 408], 6249))
print(s.cache)
