package _1_78

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	l := mergeKLists(lists[:mid])
	r := mergeKLists(lists[mid:])
	return mergeList(l, r)
}
