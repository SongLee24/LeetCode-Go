/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    point := res
    upper := int(0)
    for ; l1 != nil && l2 != nil; {
        val := (l1.Val + l2.Val + upper) % 10
        upper = (l1.Val + l2.Val + upper) / 10
        l1 = l1.Next
        l2 = l2.Next
        if l1 == nil && l2 == nil {
            if upper != 0 {
                point.Next = &ListNode{
                    Val: upper,
                    Next: nil,
                }
            } else {
                point.Next = nil
            }
        } else {
            point.Next = &ListNode{}
        } 
        point.Val = val
        point = point.Next
        
    }

    for ; l1 != nil; {
        val := int((l1.Val + upper) % 10)
        upper = int((l1.Val + upper) / 10)
        l1 = l1.Next
        if l1 == nil {
            if upper != 0 {
                point.Next = &ListNode{
                    Val: upper,
                    Next: nil,
                }
            } else {
                point.Next = nil
            }
        } else {
            point.Next = &ListNode{}
        }
        point.Val = val
        point = point.Next
        
    }

    for ; l2 != nil; {
        val := int((l2.Val + upper) % 10)
        upper = int((l2.Val + upper) / 10)
        l2 = l2.Next
        if l2 == nil {
            if upper != 0 {
                point.Next = &ListNode{
                    Val: upper,
                    Next: nil,
                }
            } else {
                point.Next = nil
            }
        } else {
            point.Next = &ListNode{}
        }
        point.Val = val
        point = point.Next
    }

    return res
}

