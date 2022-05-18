class Solution:
    def __init__(self):
        self.time = 0
        self.tin = {}
        self.up = []
        self.graph = []
        self.answer = []

    def resetValues(self, size):
        self.time = 0
        self.tin = {}
        max_val = 10 ** 5 + 1
        self.up = [max_val for i in range(size)]
        self.graph = [[] for i in range(size)]
        self.answer = []

    def construct_graph(self, size, edges):
        for edge in edges:
            self.graph[edge[0]].append(edge[1])
            self.graph[edge[1]].append(edge[0])

    def dfs(self, node, parent):
        self.tin[node] = self.time
        self.up[node] = self.time

        self.time += 1

        for next_node in self.graph[node]:
            if next_node == parent:
                continue

            if next_node in self.tin:
                self.up[node] = min(self.up[node], self.tin[next_node])
            else:
                self.dfs(next_node, node)
                self.up[node] = min(self.up[node], self.up[next_node])

            if self.tin[node] < self.up[next_node]:
                self.answer.append([node, next_node])

    def criticalConnections(self, n, connections):
        self.resetValues(n)
        self.construct_graph(n, connections)

        self.dfs(0, -1)

        return self.answer


s = Solution()
print(s.criticalConnections(10,
                            [[1, 0], [2, 0], [3, 0], [4, 1], [5, 3], [6, 1], [7, 2], [8, 1], [9, 6], [9, 3], [3, 2],
                             [4, 2], [7, 4], [6, 2], [8, 3], [4, 0], [8, 6], [6, 5], [6, 3], [7, 5], [8, 0], [8, 5],
                             [5, 4], [2, 1], [9, 5], [9, 7], [9, 4], [4, 3]]))
