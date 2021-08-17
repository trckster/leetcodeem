#include <iostream>
#include <queue>

using namespace std;

/** Definition for a binary tree node. */
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode() : val(0), left(nullptr), right(nullptr) {}

    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}

    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    bool isSameTree(TreeNode *a, TreeNode *b) {
        if (!a || !b) {
            return !a && !b;
        }

        queue < TreeNode * > aq;
        queue < TreeNode * > bq;

        aq.push(a);
        bq.push(b);

        while (!aq.empty() || !bq.empty()) {
            int na = aq.size();
            int nb = bq.size();

            if (na != nb) {
                return false;
            }


            while (na--) {
                TreeNode *nodeA = aq.front();
                TreeNode *nodeB = bq.front();

                if (nodeA->val != nodeB->val) {
                    return false;
                }

                aq.pop();
                bq.pop();

                if (!nodeA->left && nodeB->left || nodeA->left && !nodeB->left) {
                    return false;
                }

                if (!nodeA->right && nodeB->right || nodeA->right && !nodeB->right) {
                    return false;
                }

                if (nodeA->left) {
                    aq.push(nodeA->left);
                    bq.push(nodeB->left);
                }

                if (nodeA->right) {
                    aq.push(nodeA->right);
                    bq.push(nodeB->right);
                }
            }
        }

        return true;
    }
};

int main() {
    TreeNode *a = new TreeNode();
    TreeNode *b = new TreeNode();

    Solution solution;
    bool same = solution.isSameTree(a, b);

    if (same) {
        cout << "Equal trees\n";
    } else {
        cout << "Trees differ\n";
    }
}