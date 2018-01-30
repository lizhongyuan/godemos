package Node


type ListNode struct {
	// nodeData ListNodeData
	data int
	dataType string
	next *ListNode
	prev *ListNode
}

func GenIntListNode(val int) ListNode {
	node := new(ListNode)
	node.data = val
	node.dataType = "int"
	node.next = nil
	node.prev = nil
	return *node
}

//func GetListNodeData(node *ListNode) (ListNodeData, string) {
func GetListNodeData(node *ListNode) (int, string) {
	if nil == node {
		// todo handle the error
	}
	return node.data, node.dataType
}

func GetNextListNode(node *ListNode) *ListNode {
	if nil == node {
		// todo handle error
	}

	if nil == node.next {
		return nil
	} else {
		return node.next
	}
}

func GetPrevListNode(node *ListNode) *ListNode {
	if nil == node {
		// todo handle error
	}

	if nil == node.prev {
		return nil
	} else {
		return node.next
	}
}
