class Solution:
    def getTargetCopy(self, original, cloned, target):
        q1 = [original]
        q2 = [cloned]

        while len(q1) != 0:
            next_node_1 = q1[0]
            del q1[0]
            next_node_2 = q2[0]
            del q2[0]

            if next_node_1 == target:
                return next_node_2

            if next_node_1.left is not None:
                q1.append(next_node_1.left)
                q2.append(next_node_2.left)

            if next_node_1.right is not None:
                q1.append(next_node_1.right)
                q2.append(next_node_2.right)
