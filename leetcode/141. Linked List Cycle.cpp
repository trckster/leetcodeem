#include<iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;

    ListNode(int x) {
        val = x;
        next = 0;
    }
};

class Solution {
public:
    bool hasCycle(ListNode *head) {
        if (!head) {
            return false;
        }

        ListNode *slow = head, *fast = head;

        while (fast->next && fast->next->next) {
            slow = slow->next;
            fast = fast->next->next;

            if (slow == fast) {
                return true;
            }
        }

        return false;
    }
};

int main() {
    ListNode *head = new ListNode(2);
    head->next = new ListNode(5);
    head->next->next = new ListNode(3);
    head->next->next->next = head;

    Solution s;

    if (s.hasCycle(head)) {
        cout << "There is a cycle\n";
    } else {
        cout << "No cycles\n";
    }
}