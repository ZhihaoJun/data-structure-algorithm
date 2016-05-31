package bst

import (
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type Node struct {
	Val     int
	Left    *Node
	Right   *Node
	rwMutex sync.RWMutex
}

func NewFromInts(l []int) *Node {
	root := &Node{
		Val: l[0],
	}
	for i := 1; i < len(l); i++ {
		root.insert(l[i])
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

func (this *Node) Inorder(apply func(v int)) {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	this.inorder(apply)
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
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	return this.insert(v)
}

func (this *Node) insert(v int) *Node {
	if v < this.Val {
		if this.Left == nil {
			this.Left = &Node{
				Val: v,
			}
		} else {
			this.Left.insert(v)
		}
	} else {
		if this.Right == nil {
			this.Right = &Node{
				Val: v,
			}
		} else {
			this.Right.insert(v)
		}
	}
	return this
}

func (this *Node) DeleteAll(v int) *Node {
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	for this.has(v) {
		this = this.delete(v)
	}
	return this
}

/*
please use root = root.Delete(v) to make sure delete correct
*/
func (this *Node) Delete(v int) *Node {
	this.rwMutex.Lock()
	defer this.rwMutex.Unlock()
	return this.delete(v)
}

func (this *Node) deleteNode(n *Node) *Node {
	if n.Left == nil {
		return n.Right
	} else {
		l := n.Left
		rightMost := l.rightMost()
		rightMost.Right = n.Right
		return l
	}
}

func (this *Node) delete(v int) *Node {
	ret := this
	if this.Val == v {
		ret = this.deleteNode(this)
	} else if v < this.Val {
		if this.Left != nil {
			this.Left = this.Left.delete(v)
		}
	} else {
		if this.Right != nil {
			this.Right = this.Right.delete(v)
		}
	}
	return ret
}

func (this *Node) LeftMode() *Node {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	return this.leftMost()
}

func (this *Node) leftMost() *Node {
	parent := this
	l := this.Left
	for l != nil {
		parent = l
		l = l.Left
	}
	return parent
}

func (this *Node) RightMost() *Node {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	return this.rightMost()
}

func (this *Node) rightMost() *Node {
	parent := this
	r := this.Right
	for r != nil {
		parent = r
		r = r.Right
	}
	return parent
}

func (this *Node) Min() int {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	return this.leftMost().Val
}

func (this *Node) Max() int {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	return this.rightMost().Val
}

func (this *Node) SortedSlice() []int {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	var ret []int
	this.inorder(func(v int) {
		ret = append(ret, v)
	})
	return ret
}

func (this *Node) Num() int {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	var ret int64 = 0
	this.inorder(func(v int) {
		ret = atomic.AddInt64(&ret, 1)
	})
	return int(ret)
}

func (this *Node) Has(v int) bool {
	this.rwMutex.RLock()
	defer this.rwMutex.RUnlock()
	return this.has(v)
}

func (this *Node) has(v int) bool {
	if v < this.Val {
		if this.Left != nil {
			return this.Left.has(v)
		}
	} else if v == this.Val {
		return true
	} else {
		if this.Right != nil {
			return this.Right.has(v)
		}
	}
	return false
}
