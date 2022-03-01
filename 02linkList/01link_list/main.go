/*
 * @Author: shaun
 * @Date: 2022-03-01 19:13:54
 * @Last Modified by: shaun
 * @Last Modified time: 2022-03-01 19:21:43
  单链表
*/
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func ShowNode(p *Node) { // 遍历
	for p != nil {
		fmt.Println(*p)
		p = p.next // 移动指针
	}
}

func main() {
	var head = new(Node)
	head.data = 1
	var node1 = new(Node)
	node1.data = 2
	head.next = node1
	var node2 = new(Node)
	node2.data = 3
	node1.next = node2
	ShowNode(head)
}
