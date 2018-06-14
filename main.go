package main

import (
	"fmt"
	"algorithm/queue"
	"algorithm/tree"
)

func main() {

	//q := queue.Queue{}
	//q.Put("queue_1")
	//q.Put("queue_2")
	//travselQueue(&q)
	//
	//aa := q.Pop()
	//fmt.Println("aa===", aa)
	//travselQueue(&q)

	//list := linkedList.List{}
	//
	////1.往单链表末尾追加元素2, 3, 4, 5
	//list.Append(1)
	//list.Append(2)
	//list.Append(3)
	//list.Append(4)
	//
	////2.从头部添加元素head_node
	//list.Add("head_node")
	//
	//fmt.Println("长度======", list.Length())
	//
	////3.判断是否为空链表
	//bool := list.IsEmpty()
	//fmt.Println(bool)
	//
	////4.在指定位置2插入元素 2indexValue
	//list.Insert(2, "2_index_value")
	//travselLinkList(&list)
	//
	////5.是否包含元素2_index_value
	//isContain := list.Contain("2_index_value")
	//fmt.Println("isContain[2_index_value]:", isContain)
	//
	////6.删除元素2_index_value
	//list.Remove("2_index_value")
	//travselLinkList(&list)
	//
	////7.从位置2删除元素
	//list.RemoveAtIndex(2)
	//travselLinkList(&list)
	//
	//
	//fmt.Println("+++++++++++++++++++")


	//var s []int
	//s = append(s, 1)
	//s = append(s, 2)
	//s = append(s, 3)
	//fmt.Println(s)
	//
	//s = s[2: len(s)]
	//fmt.Println(s)

	tree := tree.Tree{}
	tree.Add(0)
	tree.Add(1)
	tree.Add(2)
	tree.Add(3)
	tree.Add(4)
	tree.Add(5)
	tree.Add(6)
	tree.Add(7)
	tree.Add(8)
	tree.Add(9)

	//广度优先遍历
	//tree.BreadthTravel()
	//fmt.Println("")

	//深度优先 先序遍历
	tree.PreOrder(tree.RootNode)
	fmt.Println("")

	//中序遍历
	tree.InOrder(tree.RootNode)
	fmt.Println("")

	//后序遍历
	tree.PostOrder(tree.RootNode)
	fmt.Println("")
}

func travselQueue(q *queue.Queue)  {
	//fmt.Println("-------queue----begin-------------")
	//遍历
	//head := q.GetHeadNode()
	//for head != nil {
	//	fmt.Println(head.Data)
	//	head = head.Next
	//}
	//fmt.Println("-------queue------end-------")

	var s []int
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println(s)

	s = s[1: len(s) -1]
	fmt.Println(s)

}


//func travselLinkList(list *linkedList.List) {
//	//遍历
//	head := list.GetHeadNode()
//	for head != nil {
//		fmt.Println(head.Data)
//		head = head.Next
//	}
//	fmt.Println("--------------------")
//}
