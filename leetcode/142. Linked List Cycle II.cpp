class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        if (!head) {
            return 0;
        }

        ListNode *slow = head;
        ListNode *fast = head;

        while (fast->next && fast->next->next) {
            fast = fast->next->next;
            slow = slow->next;

            if (fast == slow) {
                break;
            }
        }

        if (!fast->next || !fast->next->next) {
            return 0;
        }

        slow = head;
        while (fast != slow) {
            fast = fast->next;
            slow = slow->next;
        }

        return slow;
    }
};
