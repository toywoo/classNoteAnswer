package thirdproblem

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func insertHead(head **Node, value int) { // 이중포인터인 이유 매개변수는 복사임!
	newNode := new(Node)

	newNode.value = value
	newNode.next = (*head).next
	(*head).next = newNode
}

func insertNth(head **Node, value int, n int) {
	var i int = 0
	var p *Node = (*head).next
	n -= 2

	for p != nil && i < n {
		p = p.next
		i++
	}

	newNode := new(Node)
	newNode.value = value
	newNode.next = p.next
	p.next = newNode

}

func printList(head *Node) {
	var p *Node = head.next
	for p != nil {
		fmt.Println(p.value)
		p = p.next
	}
}

func deleteNth(head **Node, n int) {
	var i int = 0
	var p *Node = (*head).next
	n -= 2

	for p != nil && i < n {
		p = p.next
		i++
	}

	p.next = p.next.next
	// p.next는 GC에 의해 메모리 해제
}

func LinkedList() {
	var head *Node = new(Node)

	insertHead(&head, 1)
	insertHead(&head, 3)

	insertNth(&head, 2, 2)

	printList(head)

	deleteNth(&head, 1)

	print("\n")
	printList(head)
}
