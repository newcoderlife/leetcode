package print_immutable_linked_list_in_reverse

type ImmutableListNode struct{}

func (this *ImmutableListNode) getNext() ImmutableListNode { return ImmutableListNode{} }

func (this *ImmutableListNode) printValue() {}

func printLinkedListInReverse(head ImmutableListNode) {
	if head == nil {
		return
	}
	printLinkedListInReverse(head.getNext())
	head.printValue()
}
