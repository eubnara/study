/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
typedef struct ListNode LISTNODE;
#include <stdlib.h>
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    int over = 0;
    LISTNODE* ret = (LISTNODE*)malloc(sizeof(LISTNODE));
    LISTNODE* cur = ret;
    int sum = l1->val + l2->val;
    cur->val = sum;
    cur->next = NULL;
    l1 = l1->next;
    l2 = l2->next;
    while(l1 != NULL && l2 != NULL) {
        sum = l1->val + l2->val;
        LISTNODE* tmp = (LISTNODE*)malloc(sizeof(LISTNODE));
        tmp->val = sum;
        tmp->next = NULL;
        cur->next = tmp;
        cur = tmp;
        l1 = l1->next;
        l2 = l2->next;
    }
    LISTNODE* remain = NULL;
    if(l1 != NULL) {
        remain = l1;
    } else if(l2 != NULL) {
        remain = l2;
    }
    while (remain != NULL) {
        LISTNODE* tmp = (LISTNODE*)malloc(sizeof(LISTNODE));
        tmp->val = remain -> val;
        tmp->next = NULL;
        cur->next = remain;
        cur = remain;
        remain = remain->next;
    }
    cur = ret;
    cur->val += over;
    over = cur->val / 10;
    cur->val %= 10;
    while(cur->next != NULL) {
        cur = cur->next;
        cur->val += over;
        over = cur->val / 10;
        cur->val %= 10;
    }
    if (over != 0) {
        LISTNODE* tmp = (LISTNODE*)malloc(sizeof(LISTNODE));
        tmp->val = over;
        tmp->next = NULL;
        cur->next = tmp;
    }
    return ret;
}
