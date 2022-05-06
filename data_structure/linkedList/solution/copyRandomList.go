package solution

import "algorithm-pattern/data_structure/linkedList"

func copyRandomList(head *linkedList.Node) *linkedList.Node {
	hap := make(map[*linkedList.Node]*linkedList.Node, 0)
	if head == nil {
		return nil
	}
	res := helperCopy(head, hap)
	return res
}

func helperCopy(head *linkedList.Node, hap map[*linkedList.Node]*linkedList.Node) *linkedList.Node {
	res := head
	head2 := head
	hap[nil] = nil
	for head != nil {
		clone := &linkedList.Node{Val: head.Val}
		clone.Next = nil
		hap[head] = clone
		head = head.Next
	}
	for head2 != nil {
		hap[head2].Next = hap[head2.Next]
		hap[head2].Random = hap[head2.Random]
		head2 = head2.Next
	}
	return hap[res]
}

func helperCopydfs(head *linkedList.Node, hap map[*linkedList.Node]*linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}
	if val, ok := hap[head]; ok {
		return val
	}
	clone := &linkedList.Node{Val: head.Val}
	hap[head] = clone
	clone.Next = helperCopydfs(head.Next, hap)
	clone.Random = helperCopydfs(head.Random, hap)
	return clone
}
