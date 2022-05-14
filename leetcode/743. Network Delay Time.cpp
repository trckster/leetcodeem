#include <bits/stdc++.h>

using namespace std;

class Solution {
private:
    vector <vector<int>> graph;
    map<pair<int, int>, int> times;
    vector<int> distances;

    void clearGraph() {
        for (auto node: graph) {
            node.clear();
        }

        graph.clear();
        times.clear();
        distances.clear();
    }

    void resizeVectors(int size) {
        clearGraph();

        graph.resize(size);
        distances.resize(size, INT32_MAX);

    }

    void initGraph(vector <vector<int>> &times, int size) {
        resizeVectors(size);

        for (auto t: times) {
            int from = t[0] - 1, to = t[1] - 1, time = t[2];
            graph[from].push_back(to);

            auto edge = make_pair(from, to);
            this->times[edge] = time;
        }
    }

    void bfs(int from) {
        priority_queue <pair<int, int>> q;
        q.push(make_pair(0, from));

        while (!q.empty()) {
            pair<int, int> nextNode = q.top();
            int distance = nextNode.first, node = nextNode.second;

            q.pop();

            for (auto linkedNode: graph[node]) {
                int distanceToLinkedNode = distance + times[make_pair(node, linkedNode)];

                if (distanceToLinkedNode < distances[linkedNode]) {
                    distances[linkedNode] = distanceToLinkedNode;
                    q.push(make_pair(distanceToLinkedNode, linkedNode));
                }
            }
        }
    }

    int getMaxDestination() {
        int answer = 0;

        for (auto d: distances) {
            answer = max(answer, d);
        }

        return answer;
    }

public:
    int networkDelayTime(vector <vector<int>> &times, int n, int k) {
        initGraph(times, n);

        distances[--k] = 0;
        bfs(k);

        int answer = getMaxDestination();

        if (answer == INT32_MAX) {
            return -1;
        }

        return answer;
    }
};

int main() {
    vector <vector<int>> times{
            {2, 1, 1},
            {2, 3, 1},
            {3, 4, 1}
    };

    Solution s;
    cout << s.networkDelayTime(times, 4, 2) << endl;
}