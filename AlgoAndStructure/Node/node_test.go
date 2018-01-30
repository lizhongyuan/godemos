package Node

import (
	"fmt"
	"testing"
)

func Test_node(t *testing.T) {
	node := new(ListNode)
	node.dataType = "int"
	node.data = 1234
	node.next = nil
	node.prev = nil

	fmt.Print(node)
}




