class Solution:

    def __init__(self):
        self.graph = []
        self.visited = set()

    def build_graph(self, isConnected):
        for city1 in range(len(isConnected)):
            self.graph.append([])
            for city2 in range(len(isConnected[city1])):
                if isConnected[city1][city2] == 1:
                    self.graph[city1].append(city2)

    def discover_province(self, city):
        queue = [city]
        self.visited.add(city)

        while queue:
            current_city = queue.pop(0)

            for new_city in self.graph[current_city]:
                if new_city in self.visited:
                    continue

                self.visited.add(new_city)
                queue.append(new_city)

    def findCircleNum(self, isConnected) -> int:
        self.build_graph(isConnected)

        provinces_count = 0

        for city in range(len(self.graph)):
            if city in self.visited:
                continue

            self.discover_province(city)
            provinces_count += 1

        return provinces_count
