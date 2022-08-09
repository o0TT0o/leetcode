# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val 
#         self.next = next 
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        headNode = None
        lastNode = headNode
        carry = 0
        while True:
            if l1 == None and l2 ==None:
                if carry == 1:
                    tmpNode = ListNode()
                    tmpNode.val = 1
                    tmpNode.next = None
                    lastNode.next = tmpNode
                return headNode
            else:
                if l1 == None:
                    v1 = 0
                    v1Next = None
                else:
                    v1 = l1.val
                    v1Next = l1.next
                    
                if l2 == None:
                    v2 = 0
                    v2Next = None
                else:
                    v2 = l2.val
                    v2Next = l2.next
                    
                tmpNode = ListNode()
                val = v1 + v2 + carry
                if val >=10:
                    tmpNode.val = val - 10
                    carry = 1
                else:
                    tmpNode.val = val
                    carry = 0
                    
                if headNode == None:
                    headNode = tmpNode
                    lastNode = headNode
                else:
                    lastNode.next = tmpNode
                    lastNode = tmpNode
                l1 = v1Next
                l2 = v2Next
                