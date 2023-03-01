class Solution:
    def edges_to_graph(self, edges, n):
        graph = []
        for i in range(n):
            graph.append([])

        for edge in edges:
            graph[edge[0]].append(edge[1])
            graph[edge[1]].append(edge[0])

        return graph

    def validPath(self, n: int, edges, source: int, destination: int) -> bool:
        if not edges:
            return True

        queue = [source]
        visited = set(queue)
        graph = self.edges_to_graph(edges, n)

        while queue:
            current_node = queue.pop(0)

            for node in graph[current_node]:
                if node == destination:
                    return True

                if node not in visited:
                    queue.append(node)
                    visited.add(node)

        return False
