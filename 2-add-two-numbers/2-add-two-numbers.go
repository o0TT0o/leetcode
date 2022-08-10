/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    v1 := 0
    v2 := 0
    v1Next  := &ListNode{}
    v2Next  := &ListNode{}
    var headNode *ListNode
    var endNode *ListNode
    for true{
        //fmt.Println(l1)
        //fmt.Println(l2)
        //fmt.Println(headNode)
        if (l1 == nil && l2 == nil ) {
            if carry == 1{
                tmpNode := &ListNode{}
                tmpNode.Val = 1
                endNode.Next = tmpNode
            }
            //fmt.Println("end!!!")
            return headNode
        }
        //fmt.Println("loop again...")
        if ( l1 == nil ) {
            v1 = 0
            v1Next = nil
        }else{
            v1 = l1.Val
            v1Next = l1.Next
        }
        
        if ( l2 == nil ) {
            v2 = 0
            v2Next = nil
            
        }else{
            v2 = l2.Val
            v2Next = l2.Next
        }    
        
        tmpNode := &ListNode{}
        val := v1 + v2 + carry
        if val >= 10{
            tmpNode.Val = val - 10
            carry = 1
            
        }else{
            tmpNode.Val = val
            carry = 0
        }
        
        
        if ( headNode == nil ) {
            headNode = tmpNode
            endNode = headNode
        }else{
            endNode.Next = tmpNode
            endNode = tmpNode
        }
        
        l1 = v1Next
        l2 = v2Next
        //fmt.Println(headNode)
    }
    return headNode
}