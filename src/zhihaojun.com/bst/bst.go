package bst

import (
	"strconv"
	"strings"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewFromInts(l []int) *Node {
	root := &Node{
		Val: l[0],
	}
	for i := 1; i < len(l); i++ {
		root.Insert(l[i])
	}
	return root
}

func (this *Node) InorderString(seperate string) string {
	ret := this.SortedSlice()
	stringSlice := make([]string, len(ret))
	for i, v := range ret {
		stringSlice[i] = strconv.Itoa(v)
	}
	return strings.Join(stringSlice, seperate)
}

func (this *Node) inorder(apply func(v int)) {
	if this.Left != nil {
		this.Left.inorder(apply)
	}
	apply(this.Val)
	if this.Right != nil {
		this.Right.inorder(apply)
	}
}

func (this *Node) Insert(v int) *Node {
	if v < this.Val {
		if this.Left == nil {
			this.Left = &Node{
				Val: v,
			}
		} else {
			this.Left.Insert(v)
		}
	} else {
		if this.Right == nil {
			this.Right = &Node{
				Val: v,
			}
		} else {
			this.Right.Insert(v)
		}
	}
	return this
}

func (this *Node) deleteNode(n *Node) *Node {
	if n.Left == nil {
		return n.Right
	} else {
		l := n.Left
		rightMost := l.RightMost()
		rightMost.Right = n.Right
		return l
	}
}

func (this *Node) Delete(v int) *Node {
	if this.Val == v {
		this = this.deleteNode(this)
	} else if v < this.Val {
		if this.Left != nil {
			this.Left = this.Left.Delete(v)
		}
	} else {
		if this.Right != nil {
			this.Right = this.Right.Delete(v)
		}
	}
	return this
}

func (this *Node) LeftMost() *Node {
	parent := this
	l := this.Left
	for l != nil {
		parent = l
		l = l.Left
	}
	return parent
}

func (this *Node) RightMost() *Node {
	parent := this
	r := this.Right
	for r != nil {
		parent = r
		r = r.Right
	}
	return parent
}

func (this *Node) Min() int {
	return this.LeftMost().Val
}

func (this *Node) Max() int {
	return this.RightMost().Val
}

func (this *Node) SortedSlice() []int {
	var ret []int
	this.inorder(func(v int) {
		ret = append(ret, v)
	})
	return ret
}
