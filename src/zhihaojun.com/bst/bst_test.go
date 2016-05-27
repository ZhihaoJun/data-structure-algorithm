package bst

import (
	"fmt"
	"testing"
)

func treeFromList(l []int) *Node {
	root := &Node{
		Val: l[0],
	}
	for i := 1; i < len(l); i++ {
		root.Insert(l[i])
	}
	return root
}

func TestInorderString(t *testing.T) {
	fmt.Println("testing inorder")
	root := &Node{
		Val: 3,
	}
	left := &Node{
		Val: 2,
	}
	right := &Node{
		Val: 4,
	}
	root.Left = left
	root.Right = right
	inorder := root.InorderString(" ")
	if inorder != "2 3 4" {
		t.Error("expect inorder 2 3 4 but get")
		t.Error(inorder)
	} else {
		fmt.Println("PASS")
	}
}

func TestInsert(t *testing.T) {
	fmt.Println("testing insert")
	root := &Node{
		Val: 20,
	}
	root.Insert(10).Insert(3).Insert(20).Insert(2).Insert(30).Insert(40)
	inorder := root.InorderString(" ")
	if inorder != "2 3 10 20 20 30 40" {
		t.Error("expect inorder 2 3 10 20 30 40 but get")
		t.Error(inorder)
	} else {
		fmt.Println("PASS")
	}
}

func TestMin(t *testing.T) {
	fmt.Println("testing min")
	root := treeFromList([]int{30, 20, 6, 7, 9, 2, 40})
	min := root.Min()
	if min != 2 {
		t.Error("expect min 2 but get")
		t.Error(min)
	} else {
		fmt.Println("PASS")
	}
}

func TestMax(t *testing.T) {
	fmt.Println("testing max")
	root := treeFromList([]int{30, 20, 6, 7, 9, 2, 40})
	max := root.Max()
	if max != 40 {
		t.Error("expect max 40 but get")
		t.Error(max)
	} else {
		fmt.Println("PASS")
	}
}

func TestSortedSlice(t *testing.T) {
	fmt.Println("testing sorted")
	root := treeFromList([]int{30, 20, 6, 7, 9, 2, 40})
	sorted := root.SortedSlice()
	if fmt.Sprint(sorted) != "[2 6 7 9 20 30 40]" {
		t.Error("expect sorted result is 2 6 7 9 20 30 40 but get")
		t.Error(sorted)
	} else {
		fmt.Println("PASS")
	}
}

func TestDelete(t *testing.T) {
	root := NewFromInts([]int{30, 20, 6, 7, 9, 2, 40})
	root.Delete(20).Delete(70).Delete(6)
	inorder := root.InorderString(" ")
	if inorder != "2 7 9 30 40" {
		t.Error("expect delete tree inorder is 2 7 9 30 40 but get")
		t.Error(inorder)
	} else {
		fmt.Println("PASS")
	}
}
